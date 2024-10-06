package internal

import (
	"log/slog"
	"os"

	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/database"
	"github.com/mauriciofsnts/bot/internal/discord"
	"github.com/mauriciofsnts/bot/internal/server"
)

func Bootstrap(cfg config.Config) {
	config.SetupLog(cfg)

	slog.Info("Starting vulcano!")
	slog.Debug("If you can see this, debug logging is enabled!", "cool", true)

	db, err := database.New(cfg)

	if err != nil {
		slog.Error("Failed to connect to database:", "err", err)
		os.Exit(1)
	}

	err = database.Migrate(db)

	if err != nil {
		slog.Error("Failed to migrate database:", "err", err)
		os.Exit(1)
	}

	go server.StartHttpServer()
	discord.Init()
}
