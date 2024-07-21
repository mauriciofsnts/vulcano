package news

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TnArticle struct {
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

func GetTnNews(page int) ([]TnArticle, error) {
	const endpoint = "https://www.tabnews.com.br/api/v1/contents?strategy=relevant"

	res, err := http.Get(fmt.Sprintf("%s&page=%d", endpoint, page))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var articles []TnArticle

	if err := json.NewDecoder(res.Body).Decode(&articles); err != nil {
		return nil, err
	}

	return articles, nil
}
