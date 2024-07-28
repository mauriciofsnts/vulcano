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
		Aliases:     []string{"sht", "st"},
		Description: "Shorten a URL",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "url",
				Description: "URL to shorten",
				Required:    true,
			},
		},
		Handler: func(ctx *ctx.Context) *discord.MessageCreate {
			args := ctx.Args

			if len(args) == 0 {
				reply := ctx.ReplyErr(errors.New("you need to specify the type of information to generate. Available types: `cpf`, `uuid`, `cnpj`"))
				return &reply
			}

			url, err := shorten.Shortner(args[0], nil)

			if err != nil {
				reply := ctx.ReplyErr(err)
				return &reply
			}

			reply := ctx.Reply("Shortened URL", url, nil)
			return &reply
		},
	})
}
