package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/discord/t"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
)

func init() {
	bot.RegisterCommand("uptime", bot.Command{
		Name:        "uptime",
		Aliases:     []string{"uptime"},
		Description: t.Translate().Commands.Uptime.Description.Str(),
		Category:    bot.CategoryBot,
		Handler: func(ctx *bot.Context) {

			uptime := time.Since(ctx.Bot.StartedAt)

			ctx.Reply(bot.ComplexMessageData{
				Embed: discord.Embed{
					Title:       t.Translate().Commands.Uptime.Title.Str(),
					Description: i18n.Replace(t.Translate().Commands.Uptime.Response.Str(), durationAsText(uptime)),
				},
			})
		},
	})
}

func durationAsText(duration time.Duration) string {
	seconds := int(duration.Seconds())
	if seconds < 60 {
		return t.Translate().Utils.LessThanAMinute.Str()
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
		stringfy(days, t.Translate().Utils.Day.Str(), t.Translate().Utils.Days.Str()),
		stringfy(hours, t.Translate().Utils.Hour.Str(), t.Translate().Utils.Hours.Str()),
		stringfy(minutes, t.Translate().Utils.Minute.Str(), t.Translate().Utils.Minutes.Str()),
	))
}

func pluralize(i int, singular, plural string) string {
	if i == 1 {
		return singular
	}
	return plural
}
