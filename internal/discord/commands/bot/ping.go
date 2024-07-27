package bot

import (
	"log/slog"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
)

func init() {
	ctx.AttachCommand("ping", ctx.Command{
		Name:        "ping",
		Aliases:     []string{"pong"},
		Description: "Just a simple hello world command",
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(ctx *ctx.Context) *discord.MessageCreate {
			reply := ctx.Build(
				ctx.Embed(
					"🏓  Pong!",
					"Hello, world!",
					[]discord.EmbedField{},
				))

			msg, err := ctx.Client.Rest().CreateMessage(snowflake.MustParse(ctx.TriggerEvent.ChannelId), reply)

			if err == nil {
				slog.Info("Message sent: %s", msg.ID)
				ctx.Client.Rest().DeleteMessage(snowflake.MustParse(ctx.TriggerEvent.ChannelId), msg.ID)
			}

			return nil
		},
	})
}
