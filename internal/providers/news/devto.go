package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

func GetDevtoArticles(page int) ([]DevtoArticle, error) {
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
