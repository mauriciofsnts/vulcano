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
		Handler: func(ctx ctx.Context) *discord.MessageCreate {
			page, err := strconv.Atoi(ctx.Args[0])

			if err != nil {
				page = 1
			}

			articles, err := news.GetNewsAPIHeadlines(page)

			if err != nil {
				reply := ctx.Response.ReplyErr(err)
				return &reply
			}

			fields := make([]discord.EmbedField, len(articles))

			var wg sync.WaitGroup

			for i, article := range articles {
				wg.Add(1)
				go func(idx int, article news.NewsApiArticle) {
					defer wg.Done()

					shortenedUrl := ""

					shortenedUrl, err := shorten.Shortner(article.Url, nil)

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

			reply := ctx.Response.Reply("NewsAPI", "Here are the latest news from NewsAPI", fields)
			return &reply
		},
	})
}
