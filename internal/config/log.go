package config

import (
	"log/slog"
	"os"
)

func ConfigLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Logger initialized")
}
