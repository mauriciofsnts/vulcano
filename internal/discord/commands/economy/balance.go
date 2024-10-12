package economy

import (
	"errors"
	"fmt"

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

			balance, err := providers.Services.GuildMember.GetBalance(guild.String(), author.String())

			if err != nil {
				errorReply := ctx.Response.ReplyErr(errors.New("an error occurred while trying to get your balance"))
				return &errorReply
			}

			reply := ctx.Response.Reply("Balance", "Your balance is: "+fmt.Sprint(balance), []discord.EmbedField{})

			return &reply
		},
	})
}
