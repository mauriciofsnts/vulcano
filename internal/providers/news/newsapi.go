package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func GetNewsAPIHeadlines() ([]NewsApiArticle, error) {
	const endpoint = "https://newsapi.org/v2/top-headlines?country=br"

	apiKey := "YOUR_API_KEY"

	res, err := http.Get(fmt.Sprintf("%s&apiKey=%s", endpoint, apiKey))

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
