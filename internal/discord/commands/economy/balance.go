package economy

import (
	"errors"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/i18n"
	"github.com/mauriciofsnts/bot/internal/providers"
)

func init() {
	ctx.RegisterCommand("balance", ctx.Command{
		Name:        "balance",
		Aliases:     []string{"balance"},
		Description: ctx.Translate().Commands.Balance.Description.Str(),
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {
			author := data.TriggerEvent.AuthorId
			guild := data.TriggerEvent.GuildId

			balance, err := providers.Services.GuildMember.GetBalance(guild.String(), author.String())

			if err != nil {
				msg := ctx.Translate().Commands.Balance.Error.Str()
				errorReply := data.Response.BuildDefaultErrorMessage(errors.New(msg))
				return &errorReply
			}

			msg := i18n.Replace(ctx.Translate().Commands.Balance.Reply.Str(), balance)
			reply := data.Response.BuildDefaultEmbedMessage("Balance", msg, []discord.EmbedField{})

			return &reply
		},
	})
}
