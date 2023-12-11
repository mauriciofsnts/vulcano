package bot

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
)

func init() {
	bot.RegisterCommand(
		"ping",
		bot.Command{
			Name:        "ping",
			Aliases:     []string{"ping"},
			Category:    bot.CategoryBot,
			Description: "Check the bot's latency to the Discord API.",
			Handler: func(ctx *bot.Context) {
				fields := []discord.EmbedField{
					{
						Name:   "API Latency",
						Value:  formatAPILatency(ctx.Bot.State.Gateway().Latency().Milliseconds()),
						Inline: true,
					},
				}

				ctx.Reply(bot.ComplexMessageData{
					Embed: discord.Embed{Title: "Pong!", Fields: fields},
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