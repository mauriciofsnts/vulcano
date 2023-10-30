package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const endpoint = "https://st.mrzt.dev/api/v2/links"

type Response struct {
	Address     string `json:"address"`
	Banned      bool   `json:"banned"`
	CreatedAt   string `json:"created_at"`
	Id          string `json:"id"`
	Link        string `json:"link"`
	Password    bool   `json:"password"`
	Target      string `json:"target"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
	VisitCount  int    `json:"visit_count"`
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

	requestBody, err := json.Marshal(map[string]string{
		"target":       url,
		"showAdvanced": "false",
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

	return response.Link, nil
}
