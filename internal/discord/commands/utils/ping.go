package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
)

func init() {
	events.Register("ping", events.CommandInfo{
		Function: func(cm events.CommandMessage) {
			cm.Ok(&discordgo.MessageEmbed{Description: cm.T.Commands.Ping.Response.Str()})
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Pong!",
		},
	})
}
