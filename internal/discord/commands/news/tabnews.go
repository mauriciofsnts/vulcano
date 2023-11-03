package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/providers/shorten"
	"github.com/pauloo27/logger"
)

type Article struct {
	Id                  string `json:"id"`
	Owner_id            string `json:"owner_id"`
	Parent_id           string `json:"parent_id"`
	Slug                string `json:"slug"`
	Title               string `json:"title"`
	Status              string `json:"status"`
	Source_url          string `json:"source_url"`
	Created_at          string `json:"created_at"`
	Updated_at          string `json:"updated_at"`
	Published_at        string `json:"published_at"`
	Deleted_at          string `json:"deleted_at"`
	Tabcoins            int16  `json:"tabcoins"`
	Owner_username      string `json:"owner_username"`
	Children_deep_count int16  `json:"children_deep_count"`
}

func getTabNews() ([]discord.EmbedField, error) {
	res, err := http.Get("https://www.tabnews.com.br/api/v1/contents?page=1&per_page=15&strategy=relevant")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var articles []Article

	if err := json.NewDecoder(res.Body).Decode(&articles); err != nil {
		return nil, err
	}

	fields := make([]discord.EmbedField, len(articles))

	var wg sync.WaitGroup

	for i, article := range articles {
		wg.Add(1)
		go func(idx int, article Article) {
			defer wg.Done()

			shortenedUrl := ""

			shortenedUrl, err := shorten.Shortner(
				fmt.Sprintf("https://www.tabnews.com.br/%s/%s", article.Owner_username, article.Slug),
				nil,
			)

			if err != nil {
				logger.Debugf("Error shortening url: %v", err)
			}

			value := fmt.Sprintf("⭐ %d · %s · %s", article.Tabcoins, article.Owner_username, shortenedUrl)

			fields[idx] = discord.EmbedField{
				Name:   article.Title,
				Value:  value,
				Inline: false,
			}

		}(i, article)
	}

	wg.Wait()

	return fields, nil
}

func init() {
	bot.RegisterCommand("tabnews", bot.Command{
		Name:    "tabnews",
		Aliases: []string{"tn", "tab"},
		Handler: func(ctx *bot.Context) discord.Embed {
			logger.Debug("TabNews command called")

			fields, err := getTabNews()

			if err != nil {
				logger.Debugf("Error getting tabnews: %s", err)
				return ctx.SuccessEmbed(discord.Embed{
					Title:       "TabNews",
					Description: "Erro ao buscar notícias do TabNews",
				})
			}

			embed := discord.Embed{
				Title:       "TabNews",
				Description: "Notícias do TabNews",
				Fields:      fields,
			}

			return ctx.SuccessEmbed(embed)
		},
	})
}
