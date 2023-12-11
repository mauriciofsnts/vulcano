package bot

import (
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
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
	bot.State.AddHandler(func(event *gateway.ReadyEvent) {
		bot.StartedAt = time.Now()
		logger.Debug("Bot is ready!")

		var discCommands []api.CreateCommandData

		for _, command := range cmnd {
			logger.Debug("Command registered:", command.Name)

			discCommands = append(discCommands, api.CreateCommandData{
				Name:        command.Name,
				Description: command.Description,
				Options:     command.Parameters,
			})

		}

		applicationId := bot.State.Ready().Application.ID
		logger.Debug("Application ID:", applicationId)

		cmds, err := bot.State.BulkOverwriteCommands(applicationId, discCommands)

		if err != nil {
			logger.Debug("Failed to register commands:", err)
		}

		logger.Debug("Commands registered:", len(cmds))

	})

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
		startTime:    trigger.EventTime,
		Bot:          *bot,
		RawArgs:      args,
		GuildID:      trigger.GuildID,
		AuthorID:     trigger.AuthorID,
		T:            *i18n.GetLanguage("pt_BR"),
		TriggerType:  trigger.Type,
		triggerEvent: trigger,
	}

	command.Handler(ctx)
}
