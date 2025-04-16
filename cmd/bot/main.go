package main

import (
	"log/slog"
	"os"

	"github.com/mauriciofsnts/bot/internal"
	"github.com/mauriciofsnts/bot/internal/config"
)

const (
	DefaultConfigPath = "./configs/config.yaml"
)

func main() {
	cfg, err := config.LoadConfigFromFile(DefaultConfigPath)

	if err != nil {
		slog.Error("Failed to load config:", "err", err)
		os.Exit(1)
	}

	internal.Bootstrap(cfg)
}
