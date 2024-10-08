package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
)

func init() {
	ctx.RegisterCommand("uptime", ctx.Command{
		Name:        "uptime",
		Aliases:     []string{"up"},
		Description: "Shows how long the bot has been online",
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(ctx ctx.Context) *discord.MessageCreate {
			uptime := time.Since(ctx.BotStartAt)
			reply := ctx.Response.Reply("🕒  Uptime", fmt.Sprintf("I've been online for %s", durationAsText(uptime)), nil)
			return &reply
		},
	})
}

func durationAsText(duration time.Duration) string {
	seconds := int(duration.Seconds())
	if seconds < 60 {
		return "less than a minute"
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
