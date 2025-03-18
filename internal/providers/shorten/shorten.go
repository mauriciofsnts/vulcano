package shorten

import (
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/mauriciofsnts/bot/internal/config"
)

type Options struct {
	KeepAliveFor *int
	Slug         string
}

type Response struct {
	Slug        string `json:"slug"`
	Domain      string `json:"domain"`
	URL         string `json:"url"`
	OriginalURL string `json:"original_url"`
	TTL         int    `json:"ttl"`
}

// Short takes a URL string and options for customizing the shortened URL,
// sends a POST request to the endpoint with the URL and options as the payload,
// and returns the shortened URL string and any error encountered.
// The options parameter allows you to specify the keep alive duration in seconds and a custom slug for the shortened URL.

func Shortner(apiKey string, endpoint string, url string, opts *Options) (string, error) {
	// thanks to pauloo27 for the shorten url service, give him a star
	// https://github.com/pauloo27/shurl
	client := resty.New()

	payload := map[string]any{
		"original_url": url,
		"ttl":          5259600,
	}

	if opts != nil {
		if opts.KeepAliveFor != nil {
			payload["ttl"] = *opts.KeepAliveFor
		}
		if opts.Slug != "" {
			payload["slug"] = opts.Slug
		}
	}

	var response Response

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Api-Key", apiKey).
		SetBody(payload).
		SetResult(&response).
		Post(endpoint)

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != 201 {
		return "", errors.New(string(resp.Body()))
	}

	return response.URL, nil
}

type URLShortener struct {
	ShortURL func(url string, opts *Options) (string, error)
}

func New(cfg config.Config) URLShortener {
	apiKey := cfg.Shortener.ApiKey
	endpoint := cfg.Shortener.Endpoint

	return URLShortener{
		ShortURL: func(url string, opts *Options) (string, error) {
			return Shortner(apiKey, endpoint, url, opts)
		},
	}
}
