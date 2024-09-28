package news

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers/news"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
)

func init() {
	ctx.AttachCommand("devto", ctx.Command{
		Name:        "devto",
		Aliases:     []string{"devto"},
		Description: "Get the latest news from the devto website",
		Handler: func(ctx ctx.Context) *discord.MessageCreate {
			articles, err := news.GetDevtoArticles(1)

			if err != nil {
				reply := ctx.Response.ReplyErr(err)
				return &reply
			}

			fields := make([]discord.EmbedField, len(articles))

			var wg sync.WaitGroup

			for i, article := range articles {
				wg.Add(1)
				go func(idx int, article news.DevtoArticle) {
					defer wg.Done()

					shortenedUrl := ""

					shortenedUrl, err := shorten.Shortner(article.URL, nil)

					if err != nil {
						slog.Debug("Error shortening url: ", "error", err)
					}

					var value string

					if len(article.Description) > 0 {
						value = fmt.Sprintf("%s\n\n%s", article.Description, shortenedUrl)
					} else {
						value = shortenedUrl
					}

					fields[idx] = discord.EmbedField{
						Name:  article.Title,
						Value: value,
					}
				}(i, article)
			}

			wg.Wait()

			reply := ctx.Response.Reply("Devto", "Here are the latest news from Devto", fields)
			return &reply
		},
	})
}
