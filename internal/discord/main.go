package discord

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord/slash"
	"github.com/pauloo27/logger"

	// register commands
	_ "github.com/mauriciofsnts/vulcano/internal/discord/slash/categories"
)

func Start() error {
	logger.Info("Starting bot...")

	dg, err := discordgo.New("Bot " + config.Vulcano.Token)

	if err != nil {
		return err
	}

	err = dg.Open()

	if err != nil {
		return err
	}

	err = slash.Start(dg)

	if err != nil {
		return err
	}

	logger.Info("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

	return nil
}
