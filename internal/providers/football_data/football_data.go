package footballdata

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/go-resty/resty/v2"
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

func GetMatches(dateFrom, dateTo string, competitions int, apiKey string) ([]Matches, error) {
	url := fmt.Sprintf("https://api.football-data.org/v4/matches/?dateFrom=%s&dateTo=%s&competitions=%d", dateFrom, dateTo, competitions)
	slog.Info("Fetching matches", "url", url)

	client := resty.New()

	var response MatchesResponse
	_, err := client.R().
		SetHeader("X-Auth-Token", apiKey).
		SetResult(&response).
		Get(url)

	if err != nil {
		slog.Error("Error fetching matches", "error", err.Error())
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
			return GetMatches(dateFrom, dateTo, competitions, cfg.FootballData.APIKey)
		},
	}
}
