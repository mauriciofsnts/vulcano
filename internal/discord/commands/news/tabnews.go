package news

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
	"github.com/mauriciofsnts/exodia/internal/providers/news"
	"github.com/mauriciofsnts/exodia/internal/providers/shorten"
	"github.com/mauriciofsnts/exodia/internal/providers/utils"
)

func init() {
	ctx.AttachCommand("tabnews", ctx.Command{
		Name:        "Tabnews",
		Aliases:     []string{"tn", "tabnews"}, //? tabnews is really necessary?
		Description: "Get the latest news from the tabnews website",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "page",
				Description: "The page number you want to see",
				Required:    false,
				MinValue:    utils.PtrTo(1),
				MaxValue:    utils.PtrTo(99),
			},
		},
		Handler: func(ctx *ctx.Context) discord.MessageCreate {
			tnArticles, err := news.GetTnNews(1, 15)

			if err != nil {
				embed := ctx.ErrorEmbed(err)
				return ctx.Reply(embed)
			}

			fields := make([]discord.EmbedField, len(tnArticles))
			var wg sync.WaitGroup

			for i, article := range tnArticles {
				wg.Add(1)
				go func(idx int, article news.TnArticle) {
					defer wg.Done()

					shortenedUrl := ""
					shortenedUrl, err := shorten.Shortner(
						fmt.Sprintf("https://www.tabnews.com.br/%s/%s", article.Owner_username, article.Slug),
						nil,
					)

					if err != nil {
						slog.Error("Error shortening url: ", err.Error(), "")
					}

					value := fmt.Sprintf("⭐ %d · %s · %s", article.Tabcoins, article.Owner_username, shortenedUrl)

					fields[idx] = discord.EmbedField{
						Name:  article.Title,
						Value: value,
					}

				}(i, article)
			}

			wg.Wait()

			embed := ctx.Embed(
				"Latest news from Tabnews",
				"Here are the latest news from the tabnews website",
				fields,
			)

			return ctx.Reply(embed)
		},
	})
}
