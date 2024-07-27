package shorten

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/mauriciofsnts/exodia/internal/config"
)

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

// Shortner takes a URL string and an optional keepAliveFor duration in seconds,
// sends a POST request to the endpoint with the URL as the target,
// and returns the shortened URL string and any error encountered.
func Shortner(url string, keepAliveFor *int) (string, error) {
	// thanks to pauloo27 for the shorten url service, give him a star
	// https://github.com/pauloo27/shurl
	endpoint := config.Envs.Shortener.Endpoint
	apiKey := config.Envs.Shortener.ApiKey

	requestBody, err := json.Marshal(map[string]any{
		"original_url": url,
		"ttl":          5259600,
	})

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
