package bot

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
)

func init() {
	ctx.AttachCommand("services", ctx.Command{
		Name:        "services",
		Aliases:     []string{"services"},
		Description: "List of services available",
		Options:     []discord.ApplicationCommandOption{},
		Handler: func(ctx ctx.Context) *discord.MessageCreate {

			fields := []discord.EmbedField{
				{Name: "Squarefox", Value: "https://squarefox.digital/"},
				{Name: "Website", Value: "https://mrtz.dev/"},
			}

			reply := ctx.Response.Reply("Services", "Here is a list of services available", fields)
			return &reply
		},
	})
}
