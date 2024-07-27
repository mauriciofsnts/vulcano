package discord

import (
	"log/slog"
	"strings"

	disgo "github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/exodia/internal/config"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
)

func OnMessageCreatedEvent(event *events.MessageCreate) {
	message := event.Message

	if message.Author.Bot {
		return
	}

	if !strings.HasPrefix(message.Content, config.Envs.Discord.Prefix) {
		return
	}

	msg := strings.Split(message.Content, " ")

	commandName := strings.TrimPrefix(msg[0], config.Envs.Discord.Prefix)
	found, cmd := ctx.SearchCommandByAlias(commandName)

	if !found {
		// in this case is possible because is the user that is sending the message
		return
	}

	args := msg[1:]

	slog.Debug("Args: ", slog.String("args", strings.Join(args, " ")))

	trigger := ctx.TriggerEvent{
		AuthorId:       message.Author.ID.String(),
		ChannelId:      message.ChannelID.String(),
		GuildId:        message.GuildID.String(),
		MessageId:      message.ID.String(),
		EventTimestamp: message.CreatedAt,
	}

	content := ctx.Execute(args, cmd, trigger, ctx.MESSAGE, StartedAt)
	content.MessageReference = &disgo.MessageReference{MessageID: &message.ID}

	event.Client().Rest().CreateMessage(event.ChannelID, content)
}

func OnInteractionCreatedEvent(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()

	commandName := data.CommandName()
	found, cmd := ctx.SearchCommandByAlias(commandName)

	if !found {
		slog.Error("Command not found: ", slog.String("command", commandName))
		slog.Error("The user can't dispatch an unknown slash command. Check if the command is registered")
		return
	}

	trigger := ctx.TriggerEvent{
		GuildId:        event.GuildID().String(),
		ChannelId:      event.Channel().ID().String(),
		EventTimestamp: event.CreatedAt(),
	}

	var args []string

	for _, option := range data.Options {
		args = append(args, string(option.Value))
	}

	msg := ctx.Execute(args, cmd, trigger, ctx.SLASH_COMMAND, StartedAt)
	event.CreateMessage(msg)
}

func OnReadyEvent(event *events.Ready) {
	slog.Info("Bot is ready!")
}
