package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DB        Database
	LogLevel  LogLevel
	Discord   Discord
	Shortener Shortener
	News      News
}

type Database struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
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

type LogLevel string

const (
	Info  LogLevel = "INFO"
	Debug LogLevel = "DEBUG"
	Error LogLevel = "ERROR"
)

var Envs Config

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs/")

	err := viper.ReadInConfig()

	if err != nil {
		panic("An error occurred while reading the config file: " + err.Error())
	}

	if err := viper.Unmarshal(&Envs); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		panic("Unable to decode into struct, check the config file")
	}

	validate := validator.New()
	if err := validate.Struct(&Envs); err != nil {
		log.Fatalf("Missing required attributes %v\n", err)
		panic("Missing required attributes")
	}
}
