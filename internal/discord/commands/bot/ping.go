package bot

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/i18n"
	"github.com/mauriciofsnts/bot/internal/utils"
)

func init() {
	ctx.RegisterCommand("ping", ctx.Command{
		Name:        "ping",
		Aliases:     []string{"pong"},
		Description: ctx.Translate().Commands.Ping.Description.Str(),
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {
			databaseType := data.Config.DB.Type
			shortner, err := url.Parse(data.Config.Shortener.Endpoint)

			if err != nil {
				shortner = &url.URL{}
			}

			latency := data.Client.Gateway().Latency()
			latencyMsg := i18n.Replace(ctx.Translate().Commands.Ping.Reply.Str(), formatAPILatency(latency))

			uptime := time.Since(data.BotStartAt)
			uptimeMsg := i18n.Replace(ctx.Translate().Commands.Uptime.Reply.Str(), durationAsText(uptime))

			reply := data.Response.BuildDefaultEmbedMessage(
				"🏓 Pong",
				"",
				[]discord.EmbedField{
					{Name: ctx.Translate().Global.Latency.Str(), Value: latencyMsg, Inline: utils.PtrTo(false)},
					{Name: ctx.Translate().Global.Uptime.Str(), Value: uptimeMsg, Inline: utils.PtrTo(false)},
					{Name: ctx.Translate().Global.Database.Str(), Value: string(databaseType), Inline: utils.PtrTo(false)},
					{Name: ctx.Translate().Global.Shortener.Str(), Value: string(shortner.Hostname()), Inline: utils.PtrTo(false)},
				},
			)

			return &reply
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
