package utils

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/helpers"
)

func init() {
	bot.RegisterCommand(
		"shorten",
		bot.Command{
			Name:    "shorten",
			Aliases: []string{"shorten"},
			Handler: func(ctx *bot.Context) discord.Embed {
				args := ctx.RawArgs

				if len(args) == 0 {
					return ctx.SuccessEmbed(discord.Embed{
						Title:       "Shorten",
						Description: "You need to provide a URL to shorten. Example: `!shorten https://google.com`",
					})
				}

				shortened, err := helpers.Shortner(args[0], nil)

				if err != nil {
					return ctx.SuccessEmbed(discord.Embed{
						Title:       "Shorten",
						Description: "Failed to shorten URL.",
					})
				}

				return ctx.SuccessEmbed(discord.Embed{
					Title:       "Shorten",
					Description: shortened,
				})

			},
		},
	)
}
