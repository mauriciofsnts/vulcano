package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
	"github.com/pauloo27/logger"
)

func handler(cmd events.CommandMessage) {
	cmd.Ok(&discordgo.MessageEmbed{Description: "Pong!"})
}

func init() {
	logger.Info("Registering ping command...")

	events.Register("ping", events.CommandInfo{
		Function: handler,
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Pong!",
		},
	})
}
