package shorten

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/mauriciofsnts/vulcano/internal/config"
)

const endpoint = "https://url.db.cafe/api/v1/links"

type Response struct {
	Slug        string `json:"slug"`
	Domain      string `json:"domain"`
	OriginalURL string `json:"original_url"`
	TTL         int    `json:"ttl"`
}

var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:    20,
		IdleConnTimeout: 60 * time.Second,
	},
}

// Shortner takes a URL string and an optional keepAliveFor duration in seconds,
// sends a POST request to the endpoint with the URL as the target,
// and returns the shortened URL string and any error encountered.
func Shortner(url string, keepAliveFor *int) (string, error) {

	apiKey := config.Vulcano.ShurlApiKey

	if apiKey == "" {
		slog.Error("No API key provided for shortening URL")
		return "", fmt.Errorf("no API key provided for shortening URL")
	}

	requestBody, err := json.Marshal(map[string]any{
		"original_url": url,
		"ttl":          5259600,
	})

	if err != nil {
		return "", err
	}

	// add headers to the request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var response Response

	err = json.Unmarshal(body, &response)

	if err != nil {
		return "", err
	}

	responseUrl := fmt.Sprintf("https://%s/%s", response.Domain, response.Slug)

	return responseUrl, nil
}
