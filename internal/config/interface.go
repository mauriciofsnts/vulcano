package config

import (
	"log/slog"
)

type Config struct {
	Server       Server
	DB           DatabaseConfig
	Log          LogConfig
	Discord      Discord
	Shortener    Shortener
	News         News
	FootballData FootballData
}

type LogType string

const (
	LogTypeText    LogType = "text"
	LogTypeJSON    LogType = "json"
	LogTypeColored LogType = "colored"
)

type LogConfig struct {
	Level      slog.Level
	Type       LogType
	ShowSource bool
}
type Server struct {
	Port string `validate:"required"`
}

type DatabaseType string

const (
	Postgres DatabaseType = "postgres"
	Sqlite   DatabaseType = "sqlite"
)

type DatabaseConfig struct {
	Type     DatabaseType
	Postgres PostgresConfig `yaml:"postgres"`
	Sqlite   SqliteConfig   `yaml:"sqlite"`
	Migrate  bool
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type SqliteConfig struct {
	Path string
}

type Shortener struct {
	Endpoint string `validate:"required"`
	ApiKey   string `validate:"required"`
}

type Discord struct {
	Token        string `validate:"required"`
	Prefix       string `validate:"required"`
	GuildID      string `validate:"required"`
	SyncCommands bool
}

type News struct {
	NewsapiApikey string `validate:"required" mapstructure:"newsapi_apikey"`
}

type FootballData struct {
	ApiKey string `validate:"required"`
}

var Envs Config
