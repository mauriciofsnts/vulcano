package shorten

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

	requestBody, err := json.Marshal(map[string]any{
		"original_url": url,
		"ttl":          5259600,
	})

	if err != nil {
		return "", err
	}

	res, err := client.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))

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
