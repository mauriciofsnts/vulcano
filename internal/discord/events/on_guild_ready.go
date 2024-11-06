package events

import (
	disgoEvents "github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/providers"
)

func OnGuildReady(event *disgoEvents.GuildReady) {
	providers.Services.Guild.EnsureGuildExists(event.Guild)
}
