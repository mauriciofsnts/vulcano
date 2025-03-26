package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func SetupLog(cfg Config) {
	var handler slog.Handler
	level := cfg.Logging.Level
	showSource := cfg.Logging.ShowSource

	switch cfg.Logging.Type {
	case LogFormatJSON:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     level,
			AddSource: showSource,
		})
	case LogFormatColored:
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level:      level,
			TimeFormat: time.DateTime,
			AddSource:  showSource,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     level,
			AddSource: showSource,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
