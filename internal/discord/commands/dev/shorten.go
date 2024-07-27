package dev

import (
	"errors"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
	"github.com/mauriciofsnts/exodia/internal/providers/shorten"
)

func init() {
	ctx.AttachCommand("shorten", ctx.Command{
		Name:        "shorten",
		Aliases:     []string{"st"},
		Description: "Shorten a URL",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "url",
				Description: "URL to shorten",
				Required:    true,
			},
		},
		Handler: func(ctx *ctx.Context) discord.MessageCreate {
			args := ctx.Args

			if len(args) == 0 {
				return ctx.Reply(
					ctx.ErrorEmbed(
						errors.New("you need to specify the URL to shorten"),
					))
			}

			shortened, err := shorten.Shortner(args[0], nil)

			if err != nil {
				return ctx.Reply(ctx.ErrorEmbed(err))
			}

			return ctx.Reply(ctx.Embed("Shortened URL", shortened, nil))
		},
	})
}
