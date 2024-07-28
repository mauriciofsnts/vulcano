package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mauriciofsnts/exodia/internal/config"
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

func GetNewsAPIHeadlines(page int) ([]NewsApiArticle, error) {
	const endpoint = "https://newsapi.org/v2/top-headlines?country=br&pageSize=15"

	apiKey := config.Envs.News.NewsapiApikey

	res, err := http.Get(fmt.Sprintf("%s&apiKey=%s&page=%d", endpoint, apiKey, page))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var newsAPIResponse NewsAPIResponse

	if err := json.NewDecoder(res.Body).Decode(&newsAPIResponse); err != nil {
		return nil, err
	}

	return newsAPIResponse.Articles, nil
}
