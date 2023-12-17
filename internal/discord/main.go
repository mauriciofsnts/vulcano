package discord

import (
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/pauloo27/logger"

	// register commands
	_ "github.com/mauriciofsnts/vulcano/internal/discord/commands"
)

var Bot *bot.Discord

func Start() error {
	logger.Info("Starting bot...")

	bot, err := bot.New()

	if err != nil {
		return err
	}

	Bot = bot

	return nil
}
