package events

import (
	"log/slog"
	"strings"
	"time"

	"github.com/disgoorg/disgo/bot"
	disgo "github.com/disgoorg/disgo/discord"
	disgoEvents "github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers"
)

func OnMessageCreatedEvent(event *disgoEvents.MessageCreate, client bot.Client, cfg config.Config, BotStartedAt time.Time) {
	message := event.Message

	if message.Author.Bot {
		return
	}

	if message.GuildID != nil {
		message := event.Message
		guildID := message.GuildID.String()
		authorID := message.Author.ID.String()

		err := providers.Services.GuildMember.EnsureMemberValidity(guildID, authorID)
		if err != nil {
			slog.Error("Error ensuring member validity: ", slog.Any("error", err))
		}

		if !strings.HasPrefix(message.Content, cfg.Discord.Prefix) {
			providers.Services.GuildMember.IncrementMessageCount(guildID, authorID)
			return
		}

		executeMessageCommand(event, client, cfg, true, BotStartedAt)
	} else {
		if !strings.HasPrefix(event.Message.Content, cfg.Discord.Prefix) {
			return
		}

		executeMessageCommand(event, client, cfg, false, BotStartedAt)
	}
}

func executeMessageCommand(event *disgoEvents.MessageCreate, client bot.Client, cfg config.Config, isGuild bool, BotStartedAt time.Time) {
	message := event.Message
	inputMessage := strings.Split(message.Content, " ")
	commandName := strings.TrimPrefix(inputMessage[0], cfg.Discord.Prefix)
	found, cmd := ctx.GetCommandByAlias(commandName)

	if !found {
		return
	}

	args := inputMessage[1:]

	trigger := ctx.TriggerEvent{
		AuthorId:       message.Author.ID,
		ChannelId:      message.ChannelID,
		MessageId:      message.ID,
		EventTimestamp: message.CreatedAt,
	}

	if isGuild {
		trigger.GuildId = *message.GuildID
	}

	msg := ctx.Execute(args, cmd, trigger, ctx.MESSAGE, BotStartedAt, client)

	if msg != nil {
		msg.MessageReference = &disgo.MessageReference{MessageID: &message.ID}
		event.Client().Rest().CreateMessage(event.ChannelID, *msg)

		if isGuild {
			providers.Services.GuildMember.IncrementCommandCount(message.GuildID.String(), message.Author.ID.String())
		}
	}
}
