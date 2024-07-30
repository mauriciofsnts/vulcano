package news

import (
	"fmt"
	"log/slog"
	"strconv"
	"sync"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers/news"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
	"github.com/mauriciofsnts/bot/internal/providers/utils"
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
		Handler: func(ctx *ctx.Context) *discord.MessageCreate {
			page := 1

			if len(ctx.Args) > 0 {
				value, err := strconv.Atoi(ctx.Args[0])
				if err == nil && value >= 1 {
					page = value
				}
			}

			tnArticles, err := news.GetTnNews(page, 15)

			if err != nil {
				reply := ctx.Response.ReplyErr(err)
				return &reply
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

			reply := ctx.Response.Reply("Latest news from Tabnews", "Here are the latest news from the tabnews website", fields)
			return &reply
		},
	})
}
