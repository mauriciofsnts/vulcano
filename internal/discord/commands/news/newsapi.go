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
	ctx.RegisterCommand("newsapi", ctx.Command{
		Name:        "newsapi",
		Aliases:     []string{"news"},
		Description: ctx.Translate().Commands.Newsapi.Description.Str(),
		Handler: func(context ctx.Context) *discord.MessageCreate {
			articles, err := providers.News.NewsApi(5)

			if err != nil {
				reply := context.Response.ReplyErr(err)
				return &reply
			}

			fields := make([]discord.EmbedField, len(articles))

			var wg sync.WaitGroup

			for i, article := range articles {
				wg.Add(1)
				go func(idx int, article news.NewsApiArticle) {
					defer wg.Done()

					shortenedUrl := ""
					shortenedUrl, err := providers.Shorten.ShortURL(article.Url, nil)

					if err != nil {
						slog.Error("Error shortening url: ", "error", err)
						return
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

			reply := context.Response.Reply("NewsAPI", ctx.Translate().Commands.Newsapi.Reply.Str(), fields)
			return &reply
		},
	})
}
