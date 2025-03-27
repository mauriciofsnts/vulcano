package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers"
	footballdata "github.com/mauriciofsnts/bot/internal/providers/football_data"
)

func init() {
	ctx.RegisterCommand("matches", ctx.Command{
		Name:        "matches",
		Aliases:     []string{"br", "champions"},
		Description: ctx.T().Commands.Brasileirao.Description.Str(),
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {
			today := time.Now().Format("2006-01-02")
			lastSundayOfTheWeek := time.Now().AddDate(0, 0, +6).Format("2006-01-02")

			matchesbytes, err := providers.Cache.Get(context.Background(), fmt.Sprintf("%d", providers.Football.Competitions[0].Id))
			if err != nil {
				return nil
			}

			var matches []footballdata.Matches = []footballdata.Matches{}

			err = json.Unmarshal([]byte(matchesbytes), &matches)
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
				SetDescription(ctx.T().Commands.Brasileirao.Reply.Str()).
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
