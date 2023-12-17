package bootstrap

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord"
	"github.com/mauriciofsnts/vulcano/internal/server"
	"github.com/pauloo27/logger"
)

func Start() {
	logger.HandleFatal(config.LoadConfig(), "Failed to load config, check if config.yaml exists")

	go discord.Start()
	go server.StartHttpServer()

	stop := make(chan os.Signal, 1)
	//lint:ignore SA1016 i dont know, it just works lol
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-stop
	err := discord.Bot.Close()
	if err != nil {
		logger.Error("Cannot close discord bot, but we are going to close anyway")
	}

	logger.Info("Bye bye!")
}
