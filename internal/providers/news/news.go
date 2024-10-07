package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mauriciofsnts/bot/internal/config"
)

type NewsApiSource struct {
	Id   string
	Name string
}

type NewsApiArticle struct {
	Source      NewsApiSource
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

func GetFromNewsAPI(apiKey string, page int) ([]NewsApiArticle, error) {
	const endpoint = "https://newsapi.org/v2/top-headlines?country=br&pageSize=15"

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

func GetFromTabnews(page int, maxSize int) ([]TabnewsArticle, error) {
	endpoint := fmt.Sprintf("https://www.tabnews.com.br/api/v1/contents?strategy=relevant&per_page=%d", maxSize)

	res, err := http.Get(fmt.Sprintf("%s&page=%d", endpoint, page))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var articles []TabnewsArticle

	if err := json.NewDecoder(res.Body).Decode(&articles); err != nil {
		return nil, err
	}

	return articles, nil
}

type DevtoArticle struct {
	TypeOf                 string      `json:"type_of"`
	ID                     int         `json:"id"`
	Title                  string      `json:"title"`
	Description            string      `json:"description"`
	CoverImage             string      `json:"cover_image"`
	ReadablePublishDate    string      `json:"readable_publish_date"`
	SocialImage            string      `json:"social_image"`
	TagList                []string    `json:"tag_list"`
	Tags                   string      `json:"tags"`
	Slug                   string      `json:"slug"`
	Path                   string      `json:"path"`
	URL                    string      `json:"url"`
	CanonicalURL           string      `json:"canonical_url"`
	CommentsCount          int         `json:"comments_count"`
	PositiveReactionsCount int         `json:"positive_reactions_count"`
	PublicReactionsCount   int         `json:"public_reactions_count"`
	CollectionID           interface{} `json:"collection_id"`
	CreatedAt              time.Time   `json:"created_at"`
	EditedAt               time.Time   `json:"edited_at"`
	CrosspostedAt          interface{} `json:"crossposted_at"`
	PublishedAt            time.Time   `json:"published_at"`
	LastCommentAt          time.Time   `json:"last_comment_at"`
	PublishedTimestamp     time.Time   `json:"published_timestamp"`
	ReadingTimeMinutes     int         `json:"reading_time_minutes"`
	User                   struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername string `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		WebsiteURL      string `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
	} `json:"user"`
	Organization struct {
		Name           string `json:"name"`
		Username       string `json:"username"`
		Slug           string `json:"slug"`
		ProfileImage   string `json:"profile_image"`
		ProfileImage90 string `json:"profile_image_90"`
	} `json:"organization"`
}

const PER_PAGE = 10

func GetFromDevto(page int) ([]DevtoArticle, error) {
	const endpoint = "https://dev.to/api/articles"

	res, err := http.Get(fmt.Sprintf("%s?page=%d&per_page=%d", endpoint, page, PER_PAGE))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var articles []DevtoArticle

	if err := json.NewDecoder(res.Body).Decode(&articles); err != nil {
		return nil, err
	}

	return articles, nil
}

type NewsProvider struct {
	NewsApi func(page int) ([]NewsApiArticle, error)
	Tabnews func(page int, maxSize int) ([]TabnewsArticle, error)
	Devto   func(page int) ([]DevtoArticle, error)
}

func New(cfg config.Config) NewsProvider {
	newsapiApiKey := cfg.News.NewsapiApikey

	return NewsProvider{
		NewsApi: func(page int) ([]NewsApiArticle, error) {
			return GetFromNewsAPI(newsapiApiKey, page)
		},
		Tabnews: func(page int, maxSize int) ([]TabnewsArticle, error) {
			return GetFromTabnews(page, maxSize)
		},
		Devto: func(page int) ([]DevtoArticle, error) {
			return GetFromDevto(page)
		},
	}
}
