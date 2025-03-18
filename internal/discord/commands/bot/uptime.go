package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/i18n"
)

func init() {
	ctx.RegisterCommand("uptime", ctx.Command{
		Name:        "uptime",
		Aliases:     []string{"up"},
		Description: ctx.Translate().Commands.Uptime.Description.Str(),
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(context ctx.Context) *discord.MessageCreate {
			uptime := time.Since(context.BotStartAt)
			msg := i18n.Replace(ctx.Translate().Commands.Uptime.Reply.Str(), durationAsText(uptime), "v0")
			reply := context.Response.Reply("Uptime", msg, nil)
			return &reply
		},
	})
}

func durationAsText(duration time.Duration) string {
	seconds := int(duration.Seconds())

	if seconds < 60 {
		return strings.ToLower(ctx.Translate().Global.LessThatAMinute.Str())
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
