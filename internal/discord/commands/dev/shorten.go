package dev

import (
	"errors"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
	"github.com/mauriciofsnts/bot/internal/utils"
)

func init() {
	ctx.RegisterCommand("shorten", ctx.Command{
		Name:        "shorten",
		Aliases:     []string{"sht", "st"},
		Description: ctx.Translate().Commands.Shorten.Description.Str(),
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
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {
			args := data.Args

			if len(args) == 0 {
				msg := ctx.Translate().Commands.Shorten.Error.Str()
				reply := data.Response.ReplyErr(errors.New(msg))
				return &reply
			}

			url, err := providers.Shorten.ShortURL(args[0], &shorten.Options{KeepAliveFor: utils.PtrTo(0)})

			if err != nil {
				reply := data.Response.ReplyErr(err)
				return &reply
			}

			reply := data.Response.Reply("Shortened URL", url, nil)
			return &reply
		},
	})
}
