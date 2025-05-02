package news

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/providers/cache"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
)

type NewsApiSource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NewsApiArticle struct {
	Source      NewsApiSource `json:"source"`
	Author      *string       `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	UrlToImage  string        `json:"urlToImage"`
	PublishedAt time.Time     `json:"publishedAt"`
	Content     string        `json:"content"`
}

type NewsAPIResponse struct {
	Status       string           `json:"status"`
	TotalResults int              `json:"totalResults"`
	Articles     []NewsApiArticle `json:"articles"`
}

type TabnewsArticle struct {
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

type DevtoArticle struct {
	TypeOf                 string   `json:"type_of"`
	ID                     int      `json:"id"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
	CoverImage             string   `json:"cover_image"`
	ReadablePublishDate    string   `json:"readable_publish_date"`
	SocialImage            string   `json:"social_image"`
	TagList                []string `json:"tag_list"`
	Tags                   string   `json:"tags"`
	Slug                   string   `json:"slug"`
	Path                   string   `json:"path"`
	URL                    string   `json:"url"`
	CanonicalURL           string   `json:"canonical_url"`
	CommentsCount          int      `json:"comments_count"`
	PositiveReactionsCount int      `json:"positive_reactions_count"`
	PublicReactionsCount   int      `json:"public_reactions_count"`
	ReadingTimeMinutes     int      `json:"reading_time_minutes"`
	User                   struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername string `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		WebsiteURL      string `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
	} `json:"user"`
}

var client = resty.New()

func fetchAndDecode(url string, target interface{}) error {
	resp, err := client.R().SetResult(target).Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}
	return nil
}

func GetFromNewsAPI(apiKey string, page int) ([]NewsApiArticle, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=br&pageSize=15&apiKey=%s&page=%d", apiKey, page)
	var headlines NewsAPIResponse
	if err := fetchAndDecode(url, &headlines); err != nil {
		return nil, err
	}
	return headlines.Articles, nil
}

func GetFromTabnews(page int, maxSize int) ([]TabnewsArticle, error) {
	url := fmt.Sprintf("https://www.tabnews.com.br/api/v1/contents?strategy=relevant&per_page=%d&page=%d", maxSize, page)
	var articles []TabnewsArticle
	if err := fetchAndDecode(url, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func GetTabnewsCached(c cache.Valkey, shorten shorten.URLShortener, page, maxSize int) ([]TabnewsArticle, error) {
	minPage := 1

	if page < minPage {
		page = minPage
	}

	shtUrl := func(articles []TabnewsArticle) []TabnewsArticle {
		for i, v := range articles {
			url, _ := shorten.ShortenLink(fmt.Sprintf("https://www.tabnews.com.br/%s/%s", v.Owner_username, v.Slug), nil)

			articles[i].Slug = url
		}

		return articles
	}

	key := fmt.Sprintf("tn_%d_%s5", page, time.Now().Format("02-01-2006"))
	data, err := c.Get(context.Background(), key)

	if err != nil {
		articles, err := GetFromTabnews(page, maxSize)
		articlesShortned := shtUrl(articles)

		if err == nil {
			c.Set(context.Background(), key, articlesShortned, 1*time.Hour)
		}

		return articlesShortned, err
	}

	var response []TabnewsArticle
	err = json.Unmarshal([]byte(data), &response)

	if err != nil {
		articles, err := GetFromTabnews(page, maxSize)
		articlesShortned := shtUrl(articles)

		return articlesShortned, err
	}

	return response, nil
}

const PER_PAGE = 10

func GetFromDevto(page int) ([]DevtoArticle, error) {
	url := fmt.Sprintf("https://dev.to/api/articles?page=%d&per_page=%d", page, PER_PAGE)
	var articles []DevtoArticle
	if err := fetchAndDecode(url, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

type NewsProvider struct {
	NewsApi func(page int) ([]NewsApiArticle, error)
	Tabnews func(page int, maxSize int) ([]TabnewsArticle, error)
	Devto   func(page int) ([]DevtoArticle, error)
}

func New(cfg config.Config, c cache.Valkey, st shorten.URLShortener) NewsProvider {
	return NewsProvider{
		NewsApi: func(page int) ([]NewsApiArticle, error) {
			return GetFromNewsAPI(cfg.News.APIKey, page)
		},
		Tabnews: func(page, maxSize int) ([]TabnewsArticle, error) {
			return GetTabnewsCached(c, st, page, maxSize)
		},
		Devto: GetFromDevto,
	}
}
