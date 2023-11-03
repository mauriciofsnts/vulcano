package bot

import (
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/pauloo27/logger"
)

func (bot Bot) onInteractionCreateEvent(event *gateway.InteractionCreateEvent) {
	if event.Member.User.Bot {
		return
	}

	switch data := event.Data.(type) {
	case *discord.CommandInteraction:

		cmd, found := GetCommand(data.Name)

		if !found {
			logger.Debug("Command not found")
			return
		}

		trigger := TriggerEvent{
			Type:          CommandTriggerSlash,
			EventTime:     time.Now(),
			AuthorID:      event.Sender().ID.String(),
			InteractionID: &event.InteractionEvent.ID,
			GuildID:       event.GuildID.String(),
			ChannelID:     event.ChannelID,
			Token:         event.Token,
		}

		var args []string

		for _, option := range data.Options {
			args = append(args, option.String())
		}

		handleEvent(args, cmd, &bot, trigger)
	}

}

func (bot Bot) onMessageCreateEvent(event *gateway.MessageCreateEvent) {
	if event.Author.Bot {
		return
	}

	if !strings.HasPrefix(event.Content, bot.config.Prefix) {
		return
	}

	message := strings.Split(event.Content, " ")

	commandName := strings.TrimPrefix(message[0], bot.config.Prefix)

	cmd, found := GetCommandByNameAndAliases(commandName)

	if !found {
		logger.Debug("Command not found")
		return
	}

	args := message[1:]

	trigger := TriggerEvent{
		Type:      CommandTriggerMessage,
		EventTime: time.Now(),
		AuthorID:  event.Author.ID.String(),
		GuildID:   event.GuildID.String(),
		ChannelID: event.ChannelID,
		MessageID: &event.ID,
	}

	handleEvent(args, cmd, &bot, trigger)
}

func (bot Bot) InitHandler() {
	bot.State.AddHandler(func(event *gateway.InteractionCreateEvent) {
		bot.onInteractionCreateEvent(event)
	})

	bot.State.AddHandler(func(event *gateway.MessageCreateEvent) {
		bot.onMessageCreateEvent(event)
	})
}

func handleEvent(
	args []string,
	command *Command,
	bot *Bot,
	trigger TriggerEvent,
) {

	ctx := &Context{
		startTime: trigger.EventTime,
		Bot:       *bot,
		RawArgs:   args,
		GuildID:   trigger.GuildID,
		AuthorID:  trigger.AuthorID,
	}

	var embeds []discord.Embed

	embeds = append(embeds, command.Handler(ctx))

	if trigger.Type == CommandTriggerSlash {
		bot.State.RespondInteraction(*trigger.InteractionID, trigger.Token, api.InteractionResponse{
			Type: api.MessageInteractionWithSource,
			Data: &api.InteractionResponseData{
				Embeds: &embeds,
			},
		})
	} else {
		bot.State.SendEmbedReply(trigger.ChannelID, *trigger.MessageID, embeds[0])
	}

}
