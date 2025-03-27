package config

import (
	"log/slog"
)

type Config struct {
	Server       ServerConfig
	DB           DatabaseConfig
	Logging      LoggingConfig
	Discord      DiscordConfig
	URLShortener URLShortenerConfig
	NewsAPI      NewsAPIConfig
	FootballData FootballAPIConfig
	Valkey       ValkeyConfig
}

type LogFormat string

const (
	LogFormatText    LogFormat = "text"
	LogFormatJSON    LogFormat = "json"
	LogFormatColored LogFormat = "colored"
)

type LoggingConfig struct {
	Level      slog.Level
	Type       LogFormat
	ShowSource bool
}
type ServerConfig struct {
	Port int16 `validate:"required"`
}

type DatabaseEngine string

const (
	DatabasePostgres DatabaseEngine = "postgres"
	DatabaseSQLite   DatabaseEngine = "sqlite"
)

type DatabaseConfig struct {
	Type     DatabaseEngine
	Postgres PostgresConfig `yaml:"postgres"`
	Sqlite   SQLiteConfig   `yaml:"sqlite"`
	Migrate  bool
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type SQLiteConfig struct {
	Path string
}

type URLShortenerConfig struct {
	Endpoint string `validate:"required"`
	APIKey   string `validate:"required"`
}

type DiscordConfig struct {
	Token        string `validate:"required"`
	Prefix       string `validate:"required"`
	GuildID      string `validate:"required"`
	SyncCommands bool
}

type NewsAPIConfig struct {
	APIKey string `validate:"required" mapstructure:"newsapi_apikey"`
}

type FootballAPIConfig struct {
	Seed   bool
	APIKey string `validate:"required"`
}

type ValkeyConfig struct {
	Address  string `validate:"required"`
	Password string `validate:"required"`
}

var Envs Config
