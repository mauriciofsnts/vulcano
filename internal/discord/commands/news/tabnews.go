package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
	"github.com/mauriciofsnts/vulcano/internal/helpers"
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

func getTabNews() ([]*discordgo.MessageEmbedField, error) {
	res, err := http.Get("https://www.tabnews.com.br/api/v1/contents?page=1&per_page=15&strategy=relevant")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var articles []Article

	if err := json.NewDecoder(res.Body).Decode(&articles); err != nil {
		return nil, err
	}

	fields := make([]*discordgo.MessageEmbedField, len(articles))
	fieldChan := make(chan *discordgo.MessageEmbedField)

	var wg sync.WaitGroup

	for i, article := range articles {
		wg.Add(1)
		go func(idx int, article Article) {
			defer wg.Done()

			shortenedUrl, err := helpers.Shortner(
				fmt.Sprintf("https://www.tabnews.com.br/%s/%s", article.Owner_username, article.Slug),
			)

			if err != nil {
				logger.Debugf("Error shortening url: %v", err)
				fieldChan <- &discordgo.MessageEmbedField{}
				return
			}

			logger.Debugf("Shortened url: %s", shortenedUrl)

			value := fmt.Sprintf("⭐ %d · %s · %s", article.Tabcoins, article.Owner_username, shortenedUrl)

			fieldChan <- &discordgo.MessageEmbedField{
				Name:   article.Title,
				Value:  value,
				Inline: false,
			}
		}(i, article)
	}

	go func() {
		wg.Wait()
		close(fieldChan)
	}()

	var i int
	for field := range fieldChan {
		if field.Name != "" {
			fields[i] = field
			i++
		}

		if i >= len(articles) {
			break
		}
	}

	return fields, nil
}

func init() {
	events.Register("tabnews", events.CommandInfo{
		Function: func(cm events.CommandMessage) {

			fields, err := getTabNews()
			if err != nil {
				logger.Debug("Error getting tabnews: %s", err)
				return
			}

			embed := &discordgo.MessageEmbed{
				Title:       cm.T.Commands.Tabnews.Title.Str(),
				Description: cm.T.Commands.Tabnews.Description.Str(),
				Fields:      fields,
			}

			cm.Ok(embed)
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "tabnews",
			Description: "Confira as últimas notícias do site TabNews.",
		},
	})
}
