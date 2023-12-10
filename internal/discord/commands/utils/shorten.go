package utils

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/providers/shorten"
)

func init() {
	bot.RegisterCommand(
		"shorten",
		bot.Command{
			Name:     "shorten",
			Aliases:  []string{"st"},
			Category: "🔧 Utils",
			Parameters: []discord.CommandOption{
				&discord.StringOption{
					OptionName:  "url",
					Description: "URL to shorten",
					Required:    true,
				},
			},
			Handler: func(ctx *bot.Context) {
				args := ctx.RawArgs

				if len(args) == 0 {
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       "Shorten",
							Description: "You need to provide a URL to shorten. Example: `!shorten https://google.com`",
						},
					})
					return
				}

				shortened, err := shorten.Shortner(args[0], nil)

				if err != nil {
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       "Shorten",
							Description: "Failed to shorten URL.",
						},
					})
					return
				}

				ctx.Reply(bot.ComplexMessageData{
					Embed: discord.Embed{
						Title:       "Shorten",
						Description: shortened,
					},
				})

			},
		},
	)
}
