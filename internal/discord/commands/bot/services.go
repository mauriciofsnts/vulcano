package bot

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
)

func init() {
	ctx.RegisterCommand("services", ctx.Command{
		Name:        "services",
		Aliases:     []string{"services"},
		Description: ctx.Translate().Commands.Services.Description.Str(),
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {

			fields := []discord.EmbedField{
				{Name: "Squarefox", Value: "https://squarefox.digital/"},
				{Name: "Website", Value: "https://mrtz.dev/"},
			}

			reply := data.Response.Reply("Services", string(ctx.Translate().Commands.Services.Reply.Str()), fields)
			return &reply
		},
	})
}
