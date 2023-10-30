package utils

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
	"github.com/mauriciofsnts/vulcano/internal/helpers"
	"github.com/pauloo27/logger"
)

func init() {
	events.Register("shorten", events.CommandInfo{
		Function: func(cm events.CommandMessage) {
			var argsURL string
			var argsDuration *int64

			if cm.Interaction != nil {
				argsURL = cm.Interaction.Args[0].StringValue()
				duration := cm.Interaction.Args[1].IntValue()
				argsDuration = &duration
			} else {
				argsURL = cm.Message.Args[0]
				duration, _ := strconv.ParseInt(cm.Message.Args[1], 10, 64)
				argsDuration = &duration
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
