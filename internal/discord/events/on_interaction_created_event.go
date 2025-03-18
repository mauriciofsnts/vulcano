package events

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/disgoorg/disgo/bot"
	disgoEvents "github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers"
)

func OnInteractionCreatedEvent(event *disgoEvents.ApplicationCommandInteractionCreate, client bot.Client, cfg config.Config, BotStartedAt time.Time) {
	data := event.SlashCommandInteractionData()

	commandName := data.CommandName()
	found, cmd := ctx.GetCommandByAlias(commandName)

	if !found {
		slog.Error("Command not found: ", slog.String("command", commandName))
		slog.Error("The user can't dispatch an unknown slash command. Check if the command is registered")
		return
	}

	trigger := ctx.TriggerEvent{
		ChannelId:      event.Channel().ID(),
		EventTimestamp: event.CreatedAt(),
		AuthorId:       event.User().ID,
	}

	guildId := event.GuildID()

	if guildId != nil {
		trigger.GuildId = *guildId

		err := providers.Services.GuildMember.EnsureMemberValidity(guildId.String(), event.User().ID.String())

		if err != nil {
			slog.Error("Error ensuring member validity: ", slog.Any("error", err))
		}
	}

	var args []string

	for _, option := range data.Options {
		var value any

		err := json.Unmarshal(option.Value, &value)

		if err != nil {
			return
		}

		args = append(args, fmt.Sprintf("%v", value))
	}

	msg := ctx.Execute(args, cmd, trigger, ctx.SLASH_COMMAND, BotStartedAt, client, cfg)

	if msg != nil {
		event.CreateMessage(*msg)
		providers.Services.GuildMember.IncrementCommandCount(guildId.String(), event.User().ID.String())
	}
}
