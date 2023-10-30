package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
	"github.com/mauriciofsnts/vulcano/internal/helpers"
	"github.com/pauloo27/logger"
)

func init() {
	events.Register("shorten", events.CommandInfo{
		Function: func(cm events.CommandMessage) {
			var argsDuration *int

			argsURL := cm.GetArgString(0)

			if cm.HasArg(1) {
				duration, err := cm.GetArgInt(1)

				if err == nil {
					argsDuration = &duration
				}

			}

			shortenedURL, err := helpers.Shortner(argsURL, argsDuration)

			if err != nil {
				logger.Error("Error shortening url:", err)
				return
			}

			cm.Ok(&discordgo.MessageEmbed{
				Title:       "Shortened URL",
				Description: shortenedURL,
			})

		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "shorten",
			Description: "Shorten a URL",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "url",
					Description: "URL to shorten",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "duration",
					Description: "Duration of the shortened URL",
					Required:    false,
				},
			},
		},
	})
}
