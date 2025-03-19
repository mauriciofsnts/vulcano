package news

import (
	"fmt"
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
		Description: ctx.T().Commands.Devto.Description.Str(),
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {
			articles, err := providers.News.Devto(5)

			if err != nil {
				reply := data.Response.BuildDefaultErrorMessage(err)
				return &reply
			}

			fields := make([]discord.EmbedField, len(articles))

			var wg sync.WaitGroup

			for i, article := range articles {
				wg.Add(1)
				go func(idx int, article news.DevtoArticle) {
					defer wg.Done()

					url, _ := providers.Shorten.ShortenLink(article.URL, nil)

					var value string

					if len(article.Description) > 0 {
						value = fmt.Sprintf("%s\n%s", article.Description, url)
					} else {
						value = url
					}

					fields[idx] = discord.EmbedField{
						Name:  article.Title,
						Value: value,
					}
				}(i, article)
			}

			wg.Wait()

			reply := data.Response.BuildDefaultEmbedMessage("Devto", string(ctx.T().Commands.Devto.Reply.Str()), fields)
			return &reply
		},
	})
}
