package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
)

func init() {
	bot.RegisterCommand("uptime", bot.Command{
		Name:        "uptime",
		Aliases:     []string{"uptime"},
		Description: "Shows how long the bot has been online",
		Category:    "🤖 Bot",
		Handler: func(ctx *bot.Context) {

			uptime := time.Since(ctx.Bot.StartedAt)

			ctx.Reply(bot.ComplexMessageData{
				Embed: discord.Embed{Title: "Uptime", Description: durationAsText(uptime)},
			})
		},
	})
}

func durationAsText(duration time.Duration) string {
	seconds := int(duration.Seconds())
	if seconds < 60 {
		return "Less than a minute"
	}

	days := seconds / 86400
	hours := (seconds % 86400) / 3600
	minutes := (seconds % 3600) / 60

	stringfy := func(i int, singular, plural string) string {
		if i == 0 {
			return ""
		}
		return fmt.Sprintf("%d %s", i, pluralize(i, singular, plural))
	}

	return strings.TrimSpace(fmt.Sprintf(
		"%s %s %s",
		stringfy(days, "Day", "Days"),
		stringfy(hours, "Hour", "Hours"),
		stringfy(minutes, "Minute", "Minutes"),
	))
}

func pluralize(i int, singular, plural string) string {
	if i == 1 {
		return singular
	}
	return plural
}
