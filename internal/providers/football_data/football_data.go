package footballdata

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/providers/cache"
	"github.com/mauriciofsnts/bot/internal/providers/cron"
)

const (
	footballDataBaseURL  = "https://api.football-data.org/v4/matches/"
	defaultCacheDuration = 7 * 24 * time.Hour
	cronScheduleInterval = time.Minute * 2
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

type LeagueInfo struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
}

// Predefined list of competitions
var competitions = []LeagueInfo{
	{Id: 2013, Name: "Campeonato Brasileiro Série A"},
	{Id: 2152, Name: "Copa Libertadores"},
	{Id: 2081, Name: "Copa Sudamericana"},
	{Id: 2037, Name: "Copa do Brasil"},
	{Id: 2021, Name: "Premier League"},
	{Id: 2001, Name: "UEFA Champions League"},
	{Id: 2154, Name: "UEFA Conference League"},
	{Id: 2018, Name: "European Championship"},
	{Id: 2078, Name: "Supercopa de España"},
	{Id: 2079, Name: "Copa del Rey"},
	{Id: 2015, Name: "Ligue 1"},
	{Id: 2002, Name: "Bundesliga"},
	{Id: 2019, Name: "Serie A"},
	{Id: 2014, Name: "Primera Division"},
	{Id: 2000, Name: "FIFA World Cup"},
	{Id: 2178, Name: "FIFA Club World Cup"},
	{Id: 2082, Name: "WC Qualification CONMEBOL"},
}

func GetMatches(dateFrom, dateTo string, competitions int, apiKey string) ([]Matches, error) {
	url := fmt.Sprintf("%s?dateFrom=%s&dateTo=%s&competitions=%d",
		footballDataBaseURL, dateFrom, dateTo, competitions)

	client := resty.New().
		SetTimeout(10 * time.Second).
		SetRetryCount(2).
		SetRetryWaitTime(2 * time.Second)

	var response MatchesResponse
	resp, err := client.R().
		SetHeader("X-Auth-Token", apiKey).
		SetResult(&response).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed on request: %w", err)
	}

	if resp.IsError() {
		slog.Error("failed response",
			"status", resp.Status(),
			"body", string(resp.Body()))
		return nil, fmt.Errorf("%s", resp.Status())
	}

	return response.Matches, nil
}

func fetchAndCacheMatches(
	competition LeagueInfo,
	cache cache.Valkey,
	apiKey string,
) ([]Matches, error) {
	today := time.Now().Format("2006-01-02")
	lastSundayOfWeek := time.Now().AddDate(0, 0, +6).Format("2006-01-02")

	matches, err := GetMatches(today, lastSundayOfWeek, int(competition.Id), apiKey)
	if err != nil {
		return nil, fmt.Errorf("error on fetching matches %s: %w", competition.Name, err)
	}

	err = cache.Set(
		context.Background(),
		fmt.Sprintf("%d", competition.Id),
		matches,
		defaultCacheDuration,
	)
	if err != nil {
		slog.Error("error on caching data", "competitionId", competition.Id, "error", err.Error())
		return matches, fmt.Errorf("error on caching data: %w", err)
	}

	return matches, nil
}

func cacheMatchesForCompetitions(
	competitions []LeagueInfo,
	cache cache.Valkey,
	apiKey string,
) {
	competitionChunks := slices.Chunk(competitions, 5)

	for chunk := range competitionChunks {
		for _, competition := range chunk {
			_, err := cache.Get(context.Background(), fmt.Sprintf("%d", competition.Id))
			if err != nil {
				_, cacheErr := fetchAndCacheMatches(competition, cache, apiKey)
				if cacheErr != nil {
					slog.Error("failed to search and cache matches", "competitionId", competition.Id, "error", cacheErr)
				}
			}
		}

		time.Sleep(cronScheduleInterval)
	}
}

func configureCronJobs(
	cron cron.Cron,
	cache cache.Valkey,
	cfg config.Config,
) error {
	if cfg.FootballData.APIKey == "" {
		return fmt.Errorf("missing footballdata api key")
	}

	competitions := competitions

	for i, competition := range competitions {
		cronExpression := fmt.Sprintf("%d 0 * * 6", i+1) // "At 00:01 on Saturday."

		err := cron.AddJob(cronExpression, func() {
			cacheMatchesForCompetitions([]LeagueInfo{competition}, cache, cfg.FootballData.APIKey)
		})

		if err != nil {
			slog.Error("error on add cron job", "competitionId", competition.Id, "error", err.Error())
			return err
		}
	}

	cron.Start()
	return nil
}

type FootballDataProvider struct {
	GetMatches   func(dateFrom, dateTo string, competitions int) ([]Matches, error)
	Competitions []LeagueInfo
}

func New(cfg config.Config, cron cron.Cron, cache cache.Valkey) FootballDataProvider {
	go func() {
		err := configureCronJobs(cron, cache, cfg)
		if err != nil {
			slog.Error("failed to configurate cron job", "error", err)
			os.Exit(1)
		}
	}()

	return FootballDataProvider{
		Competitions: competitions,
		GetMatches: func(dateFrom, dateTo string, competitions int) ([]Matches, error) {
			return GetMatches(dateFrom, dateTo, competitions, cfg.FootballData.APIKey)
		},
	}
}
