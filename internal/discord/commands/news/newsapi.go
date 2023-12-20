package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/discord/t"
	"github.com/mauriciofsnts/vulcano/internal/providers/shorten"
)

type Source struct {
	Id   string
	Name string
}

type NewsApiArticle struct {
	Source      Source
	Author      *string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt time.Time
	Content     string
}

type NewsAPIResponse struct {
	Status       string
	TotalResults int
	Articles     []NewsApiArticle
}

func getNewsAPI() ([]discord.EmbedField, error) {
	apiKey := config.Vulcano.NewsAPIKey

	if apiKey == "" {
		return nil, errors.New("NewsAPI key is empty")
	}

	res, err := http.Get("https://newsapi.org/v2/top-headlines?country=br&apiKey=" + apiKey)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var newsAPIResponse NewsAPIResponse

	if err := json.NewDecoder(res.Body).Decode(&newsAPIResponse); err != nil {
		return nil, err
	}

	fields := make([]discord.EmbedField, len(newsAPIResponse.Articles))

	var wg sync.WaitGroup

	for i, article := range newsAPIResponse.Articles {
		wg.Add(1)
		go func(idx int, article NewsApiArticle) {
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

	return fields, nil

}

func init() {
	bot.RegisterCommand("newsapi", bot.Command{
		Name:        "newsapi",
		Aliases:     []string{"news"},
		Description: t.Translate().Commands.NewsAPI.Description.Str(),
		Category:    bot.CategoryNews,
		Handler: func(ctx *bot.Context) {
			fields, err := getNewsAPI()

			if err != nil {
				slog.Debug("Error getting news from newsapi: ", err)

				ctx.ReplyError(bot.ComplexMessageData{
					Embed: discord.Embed{
						Title:       "Error",
						Description: "Error getting news from newsapi",
					},
				})

				return
			}

			ctx.Reply(bot.ComplexMessageData{
				Embed: discord.Embed{
					Title:       t.Translate().Commands.NewsAPI.Title.Str(),
					Description: t.Translate().Commands.NewsAPI.Description.Str(),
					Fields:      fields,
				},
			})

		},
	})

}
