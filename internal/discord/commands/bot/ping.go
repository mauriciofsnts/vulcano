package bot

import (
	"fmt"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/i18n"
)

func init() {
	ctx.RegisterCommand("ping", ctx.Command{
		Name:        "ping",
		Aliases:     []string{"pong"},
		Description: ctx.Translate().Commands.Ping.Description.Str(),
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(args ctx.Context) *discord.MessageCreate {
			// Latency returns the latency of the Gateway.
			// This is calculated by the time it takes to send a heartbeat and receive a heartbeat ack by discord.
			latency := args.Client.Gateway().Latency()

			msg := i18n.Replace(ctx.Translate().Commands.Ping.Reply.Str(), formatAPILatency(latency))

			args.Client.Rest().CreateMessage(args.TriggerEvent.ChannelId, discord.MessageCreate{
				Content: msg,
			})

			return nil
		},
	})
}

func formatAPILatency(latency time.Duration) string {
	ms := latency

	if ms < time.Millisecond {
		return fmt.Sprintf("%d", ms.Microseconds())
	} else if ms < time.Second {
		return fmt.Sprintf("%d", ms.Milliseconds())
	} else {
		return fmt.Sprintf("%.2f", ms.Seconds())
	}
}
