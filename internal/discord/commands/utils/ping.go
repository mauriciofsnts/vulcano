package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
)

func init() {
	events.Register("ping", events.CommandInfo{
		Function: func(cm events.CommandMessage) {
			latency := cm.Session.HeartbeatLatency().Milliseconds()
			ms := formatAPILatency(latency)

			cm.Ok(&discordgo.MessageEmbed{
				Title: cm.T.Commands.Ping.Response.Str(),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "API Latency",
						Value:  ms,
						Inline: true,
					},
				},
			})

		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Pong!",
		},
	})
}

func formatAPILatency(latency int64) string {
	ms := latency

	var icon string
	if ms < 50 {
		icon = "🟢"
	} else if ms < 100 {
		icon = "🟡"
	} else {
		icon = "🔴"
	}
	return fmt.Sprintf(
		"%s %d ms",
		icon,
		ms,
	)
}
