package news

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers"
	"github.com/mauriciofsnts/bot/internal/providers/news"
)

func init() {
	ctx.RegisterCommand("devto", ctx.Command{
		Name:        "devto",
		Aliases:     []string{"devto"},
		Description: ctx.Translate().Commands.Devto.Description.Str(),
		Handler: func(context ctx.Context) *discord.MessageCreate {
			articles, err := providers.News.Devto(5)

			if err != nil {
				reply := context.Response.ReplyErr(err)
				return &reply
			}

			fields := make([]discord.EmbedField, len(articles))

			var wg sync.WaitGroup

			for i, article := range articles {
				wg.Add(1)
				go func(idx int, article news.DevtoArticle) {
					defer wg.Done()

					shortenedUrl := ""
					shortenedUrl, err := providers.Shorten.ShortURL(article.URL, nil)

					if err != nil {
						slog.Debug("Error shortening url: ", "error", err)
						return
					}

					var value string

					if len(article.Description) > 0 {
						value = fmt.Sprintf("%s\n%s", article.Description, shortenedUrl)
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

			reply := context.Response.Reply("Devto", string(ctx.Translate().Commands.Devto.Reply.Str()), fields)
			return &reply
		},
	})
}
