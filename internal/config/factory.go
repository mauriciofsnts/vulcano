package config

import (
	"os"

	"github.com/ghodss/yaml"
)

func LoadConfigFromFile(configPath string) (Config, error) {
	/* #nosec G304 */
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	return LoadConfigFromData(data)
}

func LoadConfigFromData(data []byte) (Config, error) {
	var config Config

	err := yaml.Unmarshal(data, &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
