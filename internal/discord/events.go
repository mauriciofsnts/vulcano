package discord

import (
	"log/slog"
	"strings"

	"github.com/disgoorg/disgo/bot"
	disgo "github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	customEvents "github.com/mauriciofsnts/bot/internal/discord/events"
)

func OnMessageCreatedEvent(event *events.MessageCreate, client *bot.Client) {
	message := event.Message

	if message.Author.Bot {
		return
	}

	if !strings.HasPrefix(message.Content, config.Envs.Discord.Prefix) {
		return
	}

	inputMessage := strings.Split(message.Content, " ")

	commandName := strings.TrimPrefix(inputMessage[0], config.Envs.Discord.Prefix)
	found, cmd := ctx.SearchCommandByAlias(commandName)

	if !found {
		// in this case is possible because is the user that is sending the message
		return
	}

	args := inputMessage[1:]

	slog.Debug("Args: ", slog.String("args", strings.Join(args, " ")))

	trigger := ctx.TriggerEvent{
		AuthorId:       message.Author.ID,
		ChannelId:      message.ChannelID,
		GuildId:        *message.GuildID,
		MessageId:      message.ID,
		EventTimestamp: message.CreatedAt,
	}

	msg := ctx.Execute(args, cmd, trigger, ctx.MESSAGE, StartedAt, *client)

	if msg != nil {
		msg.MessageReference = &disgo.MessageReference{MessageID: &message.ID}
		event.Client().Rest().CreateMessage(event.ChannelID, *msg)
	}
}

func OnInteractionCreatedEvent(event *events.ApplicationCommandInteractionCreate, client *bot.Client) {
	data := event.SlashCommandInteractionData()

	commandName := data.CommandName()
	found, cmd := ctx.SearchCommandByAlias(commandName)

	if !found {
		slog.Error("Command not found: ", slog.String("command", commandName))
		slog.Error("The user can't dispatch an unknown slash command. Check if the command is registered")
		return
	}

	trigger := ctx.TriggerEvent{
		GuildId:        *event.GuildID(),
		ChannelId:      event.Channel().ID(),
		EventTimestamp: event.CreatedAt(),
		AuthorId:       event.User().ID,
	}

	var args []string

	for _, option := range data.Options {
		args = append(args, string(option.Value))
	}

	msg := ctx.Execute(args, cmd, trigger, ctx.SLASH_COMMAND, StartedAt, *client)

	if msg != nil {
		event.CreateMessage(*msg)
	}
}

func OnReadyEvent(event *events.Ready) {
	slog.Info("Bot is ready!")
}

func OnGuildChannelCreatedEvent(event *events.GuildChannelCreate, client *bot.Client) {
	channelId := event.ChannelID
	message := disgo.NewMessageCreateBuilder().SetContent("first!").Build()
	event.Client().Rest().CreateMessage(channelId, message)
}

func OnMessageReactionAddedEvent(event *events.MessageReactionAdd, client *bot.Client) {
	customEvents.OnGamble(event, *client)
}
