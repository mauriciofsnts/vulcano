package shorten

import (
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/mauriciofsnts/bot/internal/config"
)

type ShurlOptions struct {
	KeepAliveFor *int
	Slug         string
}

type ShurlResponse struct {
	Slug        string `json:"slug"`
	Domain      string `json:"domain"`
	URL         string `json:"url"`
	OriginalURL string `json:"original_url"`
	TTL         int    `json:"ttl"`
}

// ShurlShortner attempts to generate a shortened URL for the provided 'url' using the Shurl service.
// It returns the shortened URL if successful, or an error if a critical issue occurs.
// If the Shurl service fails to shorten the URL (e.g., due to an invalid request or service error), it returns an error.

// https://github.com/pauloo27/shurl
func ShurlShortner(apiKey string, endpoint string, url string, options *ShurlOptions) (string, error) {
	client := resty.New()

	payload := map[string]any{
		"original_url": url,
		"ttl":          5259600,
	}

	if options != nil {
		if options.KeepAliveFor != nil {
			payload["ttl"] = *options.KeepAliveFor
		}
		if options.Slug != "" {
			payload["slug"] = options.Slug
		}
	}

	var response ShurlResponse

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
	// ShortenLink attempts to generate a shortened URL for the provided 'url'.
	//
	// If the shortener service is unable to generate a shortened URL (e.g., due to an error or service unavailability),
	// it should gracefully return the original 'url' without throwing an error.
	//
	// Parameters:
	//   - url: The original URL to be shortened.
	//   - opts: Optional parameters for the shortening process (e.g., custom aliases, expiration).
	//
	// Returns:
	//   - string: The shortened URL, or the original URL if shortening fails.
	//   - error: An error if a critical issue occurs during the shortening process (e.g., network error, invalid input).
	//            A nil error indicates successful processing, even if the original URL was returned.
	//
	// Example:
	//   shortenedURL, err := shortener.ShortenLink("https://www.example.com", nil)
	//   if err != nil {
	//       // Handle critical errors, but not cases where the original URL was returned.
	//       log.Println("Error shortening URL:", err)
	//   }
	//   log.Println("Shortened URL:", shortenedURL)
	ShortenLink func(url string, opts *ShurlOptions) (string, error)
}

func New(cfg config.Config) URLShortener {
	apiKey := cfg.Shortener.ApiKey
	endpoint := cfg.Shortener.Endpoint

	return URLShortener{
		ShortenLink: func(url string, opts *ShurlOptions) (string, error) {
			shortenedURL, err := ShurlShortner(apiKey, endpoint, url, opts)

			if err != nil {
				return url, err
			}
			if shortenedURL == "" {
				return url, nil
			}
			return shortenedURL, nil
		},
	}
}
