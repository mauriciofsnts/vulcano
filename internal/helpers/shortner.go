package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const endpoint = "https://st.mrzt.dev/api/v2/links"

type Response struct {
	Address     string
	Banned      bool
	Created_at  string
	Id          string
	Link        string
	Password    bool
	Target      string
	Description string
	Updated_at  string
	Visit_count int
}

func Shortner(url string) (string, error) {

	requestBody, err := json.Marshal(map[string]string{
		"target":       url,
		"showAdvanced": "false",
	})

	if err != nil {
		return "", err
	}

	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))

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
