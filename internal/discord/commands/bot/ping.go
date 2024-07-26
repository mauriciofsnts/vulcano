package bot

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
)

func init() {
	ctx.AttachCommand("ping", ctx.Command{
		Name:        "ping",
		Aliases:     []string{"pong"},
		Description: "Just a simple hello world command",
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(ctx *ctx.Context) discord.MessageCreate {
			msg := discord.NewMessageCreateBuilder().SetContent("Pong!").Build()
			return msg
		},
	})
}
