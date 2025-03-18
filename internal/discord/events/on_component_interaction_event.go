package events

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	disgoEvents "github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake/v2"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
)

func OnComponentInteractionEvent(event *disgoEvents.ComponentInteractionCreate, client bot.Client) {
	found, componentState := ctx.LoadComponentStateFromDB(event.Message.ID.String())

	if !found {
		slog.Error("Button state not found: ", slog.String("message id", event.Message.ID.String()))
		return
	}

	trigger := ctx.TriggerEvent{
		AuthorId:       snowflake.MustParse(componentState.AuthorID),
		ChannelId:      snowflake.MustParse(componentState.ChannelID),
		GuildId:        snowflake.MustParse(componentState.GuildID),
		MessageId:      snowflake.MustParse(componentState.MessageID),
		EventTimestamp: event.CreatedAt(),
	}

	state := ctx.DiscordComponentContext{
		TriggerEvent: trigger,
		Client:       client,
		State:        componentState.State,
	}

	found, handler := ctx.GetComponentHandlerByName(componentState.Command)

	if !found {
		slog.Error("Component handler not found: ", slog.String("command", componentState.Command))
		return
	}

	handler(event, &state)
}
