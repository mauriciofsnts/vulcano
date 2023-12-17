package discord

import (
	"log/slog"

	"github.com/mauriciofsnts/vulcano/internal/discord/bot"

	// register commands
	_ "github.com/mauriciofsnts/vulcano/internal/discord/commands"
)

var Bot *bot.Discord

func Start() error {
	slog.Info("Starting bot...")

	bot, err := bot.New()

	if err != nil {
		return err
	}

	Bot = bot

	return nil
}
