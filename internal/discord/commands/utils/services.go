package utils

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
)

func init() {
	bot.RegisterCommand(
		"services",
		bot.Command{
			Name:        "services",
			Aliases:     []string{"services"},
			Description: "List of open services",
			Category:    "🔧 Utils",
			Handler: func(ctx *bot.Context) {

				ctx.Reply(bot.ComplexMessageData{
					Embed: discord.Embed{
						Title: "Open Services",
						URL:   "https://github.com/mauriciofsnts",
						Fields: []discord.EmbedField{
							{
								Name:   "URL Shortener",
								Value:  "https://st.mrzt.dev/",
								Inline: false,
							},
							{
								Name:   "Uptime",
								Value:  "https://uptime.mrzt.dev/",
								Inline: false,
							},
							{
								Name:   "Web Development",
								Value:  "https://speedify.dev/",
								Inline: false,
							},
						},
					},
				})

			},
		},
	)
}
