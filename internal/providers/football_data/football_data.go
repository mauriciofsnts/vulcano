package footballdata

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/mauriciofsnts/bot/internal/config"
)

type Matches struct {
	Area struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
		Flag string `json:"flag"`
	} `json:"area"`
	Competition struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Code   string `json:"code"`
		Type   string `json:"type"`
		Emblem string `json:"emblem"`
	} `json:"competition"`
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"season"`
	ID          int         `json:"id"`
	UtcDate     time.Time   `json:"utcDate"`
	Status      string      `json:"status"`
	Matchday    int         `json:"matchday"`
	Stage       string      `json:"stage"`
	Group       interface{} `json:"group"`
	LastUpdated time.Time   `json:"lastUpdated"`
	HomeTeam    struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
		Tla       string `json:"tla"`
		Crest     string `json:"crest"`
	} `json:"homeTeam"`
	AwayTeam struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
		Tla       string `json:"tla"`
		Crest     string `json:"crest"`
	} `json:"awayTeam"`
	Score struct {
		Winner   interface{} `json:"winner"`
		Duration string      `json:"duration"`
		FullTime struct {
			Home interface{} `json:"home"`
			Away interface{} `json:"away"`
		} `json:"fullTime"`
		HalfTime struct {
			Home interface{} `json:"home"`
			Away interface{} `json:"away"`
		} `json:"halfTime"`
	} `json:"score"`
	Odds struct {
		Msg string `json:"msg"`
	} `json:"odds"`
	Referees []interface{} `json:"referees"`
}

type MatchesResponse struct {
	Filters struct {
		DateFrom     string `json:"dateFrom"`
		DateTo       string `json:"dateTo"`
		Permission   string `json:"permission"`
		Competitions string `json:"competitions"`
	} `json:"filters"`
	ResultSet struct {
		Count        int    `json:"count"`
		Competitions string `json:"competitions"`
		First        string `json:"first"`
		Last         string `json:"last"`
		Played       int    `json:"played"`
	} `json:"resultSet"`
	Matches []Matches `json:"matches"`
}

// Helper function to make HTTP requests and parse the response
func fetchAndDecode(url string, target interface{}, apiKey string) error {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("X-Auth-Token", apiKey)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(target)
}

func GetMatches(dateFrom string, dateTo string, competitions int, apiKey string) ([]Matches, error) {
	url := fmt.Sprintf("https://api.football-data.org/v4/matches/?dateFrom=%s&dateTo=%s&competitions=%d", dateFrom, dateTo, competitions)
	// /v4/matches/?dateFrom=2024-11-03&dateTo=2024-11-09&competitions=2013
	slog.Info("Fetching matches", "url", url)

	var response MatchesResponse

	err := fetchAndDecode(url, &response, apiKey)

	if err != nil {
		slog.Info("Error fetching matches", "error", err.Error())
		return nil, err
	}

	return response.Matches, nil
}

type FootballDataProvider struct {
	GetMatches func(dateFrom, dateTo string, competitions int) ([]Matches, error)
}

func New(cfg config.Config) FootballDataProvider {
	return FootballDataProvider{
		GetMatches: func(dateFrom, dateTo string, competitions int) ([]Matches, error) {
			return GetMatches(dateFrom, dateTo, competitions, cfg.FootballData.ApiKey)
		},
	}
}
