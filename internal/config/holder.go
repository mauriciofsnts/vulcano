package config

import (
	"github.com/spf13/viper"
)

var (
	Vulcano *Config
)

func LoadConfig() error {
	viper.AddConfigPath(".")

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Vulcano)
	return nil
}
