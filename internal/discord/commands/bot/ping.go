package bot

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/discord/t"
)

func init() {
	bot.RegisterCommand(
		"ping",
		bot.Command{
			Name:        "ping",
			Aliases:     []string{"ping"},
			Category:    bot.CategoryBot,
			Description: t.Translate().Commands.Ping.Description.Str(),
			Handler: func(ctx *bot.Context) {
				fields := []discord.EmbedField{
					{
						Name:   "API Latency",
						Value:  formatAPILatency(ctx.Bot.State.Gateway().Latency().Milliseconds()),
						Inline: true,
					},
				}

				ctx.Reply(bot.ComplexMessageData{
					Embed: discord.Embed{Title: t.Translate().Commands.Ping.Response.Str(), Fields: fields},
				})
			}})
}

func formatAPILatency(latency int64) string {
	ms := latency

	var icon string
	if ms < 50 {
		icon = "🟢"
	} else if ms < 100 {
		icon = "🟡"
	} else {
		icon = "🔴"
	}
	return fmt.Sprintf("%s %d ms", icon, ms)
}
