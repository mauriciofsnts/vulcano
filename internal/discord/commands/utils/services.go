package utils

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/discord/t"
)

func init() {
	bot.RegisterCommand(
		"services",
		bot.Command{
			Name:        "services",
			Aliases:     []string{"services"},
			Description: t.Translate().Commands.OpenSourcesProjects.Description.Str(),
			Category:    bot.CategoryUtils,
			Handler: func(ctx *bot.Context) {

				ctx.Reply(bot.ComplexMessageData{
					Embed: discord.Embed{
						Title:       t.Translate().Commands.OpenSourcesProjects.Title.Str(),
						Description: t.Translate().Commands.OpenSourcesProjects.Description.Str(),
						URL:         "https://github.com/mauriciofsnts",
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
								Name:   "It's me",
								Value:  "https://mrtz.dev/",
								Inline: false,
							},
							{
								Name:   "SquareFox Digital",
								Value:  "https://squarefox.digital/",
								Inline: false,
							},
						},
					},
				})

			},
		},
	)
}
