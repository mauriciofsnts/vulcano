package events

import (
	"log/slog"

	disgoEvents "github.com/disgoorg/disgo/events"
)

func OnReadyEvent(event *disgoEvents.Ready) {
	slog.Info("Bot is ready!")
}
