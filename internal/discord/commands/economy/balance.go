package economy

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers"
)

func init() {
	ctx.RegisterCommand("balance", ctx.Command{
		Name:        "balance",
		Aliases:     []string{"balance"},
		Description: "Check your balance",
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(ctx ctx.Context) *discord.MessageCreate {
			author := ctx.TriggerEvent.AuthorId
			guild := ctx.TriggerEvent.GuildId

			slog.Info("Getting balance for user: ", author)
			slog.Info("Getting balance for guild: ", guild)

			balance, err := providers.Services.GuildMember.GetBalance(guild.String(), author.String())

			if err != nil {
				slog.Error("an error occurred while trying to get your balance", err)
				errorReply := ctx.Response.ReplyErr(errors.New("an error occurred while trying to get your balance"))
				return &errorReply
			}

			reply := ctx.Response.Reply("Balance", "Your balance is: "+fmt.Sprint(balance), []discord.EmbedField{})

			return &reply
		},
	})
}
