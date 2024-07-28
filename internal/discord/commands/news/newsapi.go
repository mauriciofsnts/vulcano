package news

import (
	"fmt"
	"log/slog"
	"strconv"
	"sync"

	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
	"github.com/mauriciofsnts/exodia/internal/providers/news"
	"github.com/mauriciofsnts/exodia/internal/providers/shorten"
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
			page := 1

			if len(ctx.Args) > 0 {
				value, err := strconv.Atoi(ctx.Args[0])
				if err == nil && value >= 1 {
					page = value
				}
			}

			articles, err := news.GetNewsAPIHeadlines(page)

			if err != nil {
				embed := ctx.ErrorEmbed(err)
				reply := ctx.Build(embed)
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
						slog.Debug("Error shortening url: ", err)
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

			embed := ctx.Embed("NewsAPI", "Here are the latest news from NewsAPI", fields)
			reply := ctx.Build(embed)
			return &reply
		},
	})
}
