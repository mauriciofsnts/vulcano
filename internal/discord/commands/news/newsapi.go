package news

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
	"github.com/mauriciofsnts/exodia/internal/providers/utils"
)

func init() {
	ctx.AttachCommand("newsapi", ctx.Command{
		Name:        "Newsapi",
		Aliases:     []string{"news"},
		Description: "Get the latest news from the newsapi website",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "page",
				Description: "The page number you want to see",
				Required:    false,
				MinValue:    utils.PtrTo(1),
				MaxValue:    utils.PtrTo(99),
			},
		},
		Handler: func(ctx *ctx.Context) *discord.MessageCreate {
			reply := ctx.Build(ctx.Embed("Newsapi", "This command is not implemented yet", nil))
			return &reply
		},
	})
}
