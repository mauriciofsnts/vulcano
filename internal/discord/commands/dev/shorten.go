package dev

import (
	"errors"
	"log/slog"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
	"github.com/mauriciofsnts/bot/internal/providers/utils"
)

func init() {
	ctx.RegisterCommand("shorten", ctx.Command{
		Name:        "shorten",
		Aliases:     []string{"sht", "st"},
		Description: "Shorten a URL",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "url",
				Description: "URL to shorten",
				Required:    true,
			},
			discord.ApplicationCommandOptionString{
				Name:        "slug",
				Description: "Custom alias for the shortened URL",
				Required:    false,
			},
		},
		Handler: func(ctx ctx.Context) *discord.MessageCreate {
			args := ctx.Args

			slog.Info("shorten", "args", args[1])

			if len(args) == 0 {
				reply := ctx.Response.ReplyErr(errors.New("you need to specify the type of information to generate. Available types: `cpf`, `uuid`, `cnpj`"))
				return &reply
			}

			url, err := shorten.Shortner(args[0], &shorten.Options{Slug: args[1], KeepAliveFor: utils.PtrTo(0)})

			if err != nil {
				reply := ctx.Response.ReplyErr(err)
				return &reply
			}

			reply := ctx.Response.Reply("Shortened URL", url, nil)
			return &reply
		},
	})
}
