package shorten

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

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

var httpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:    20,
		IdleConnTimeout: 60 * time.Second,
	},
}

// Short takes a URL string and options for customizing the shortened URL,
// sends a POST request to the endpoint with the URL and options as the payload,
// and returns the shortened URL string and any error encountered.
// The options parameter allows you to specify the keep alive duration in seconds and a custom slug for the shortened URL.
func Shortner(apiKey string, endpoint string, url string, opts *Options) (string, error) {
	// thanks to pauloo27 for the shorten url service, give him a star
	// https://github.com/pauloo27/shurl

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

	requestBody, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	// set the default header to perform the request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", apiKey)

	res, err := httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var response Response

	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.URL, nil
}

type URLShortener struct {
	Shorten func(url string, opts *Options) (string, error)
}

func New(cfg config.Config) URLShortener {
	apiKey := cfg.Shortener.ApiKey
	endpoint := cfg.Shortener.Endpoint

	return URLShortener{
		Shorten: func(url string, opts *Options) (string, error) {
			return Shortner(apiKey, endpoint, url, opts)
		},
	}
}
