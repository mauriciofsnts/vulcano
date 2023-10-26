package discord

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/pauloo27/logger"

	// register commands
	_ "github.com/mauriciofsnts/vulcano/internal/discord/commands"
)

func Start() error {
	logger.Info("Starting bot...")

	bot, err := bot.New()

	if err != nil {
		return err
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()

	return nil
}
