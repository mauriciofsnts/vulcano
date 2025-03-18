package bot

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers"
)

func init() {
	ctx.RegisterCommand("brasileirao", ctx.Command{
		Name:        "brasileirao",
		Aliases:     []string{"brasileirao", "br"},
		Description: ctx.Translate().Commands.Brasileirao.Description.Str(),
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {
			today := time.Now().Format("2006-01-02")
			lastSundayOfTheWeek := time.Now().AddDate(0, 0, +6).Format("2006-01-02")

			matches, err := providers.Football.GetMatches(today, lastSundayOfTheWeek, 2013)

			if err != nil {
				return nil
			}

			fields := make([]discord.EmbedField, len(matches))

			for i, match := range matches {
				parsedUtcDate := match.UtcDate

				fields[i] = discord.EmbedField{
					Value: parsedUtcDate.Format("2006-01-02"),
					Name:  fmt.Sprintf("%s x %s", match.HomeTeam.Name, match.AwayTeam.Name),
				}
			}

			messageBuilder := discord.NewMessageCreateBuilder()
			embedBuilder := discord.NewEmbedBuilder().
				SetTitle(fmt.Sprintf("Brasileirão %s / %s", today, lastSundayOfTheWeek)).
				SetDescription(ctx.Translate().Commands.Brasileirao.Reply.Str()).
				SetColor(0xffffff).
				SetFields(fields...)
			embed := embedBuilder.Build()
			messageBuilder.SetEmbeds(embed)
			msg := messageBuilder.Build()

			_, err = data.Client.Rest().CreateMessage(data.TriggerEvent.ChannelId, msg)

			if err != nil {
				slog.Error("Error creating message", "err", err.Error())
				return nil
			}

			return nil
		},
	})
}
