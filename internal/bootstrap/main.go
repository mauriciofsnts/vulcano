package bootstrap

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord"
	"github.com/mauriciofsnts/vulcano/internal/server"
)

func Start() {
	SetupLog()

	slog.Info("Starting Vulcano...")

	err := config.LoadConfig()

	if err != nil {
		slog.Error("Cannot load config file: ", err)
		os.Exit(1)
	}

	go discord.Start()

	if config.Vulcano.Port != "" {
		go server.StartHttpServer()
	}

	stop := make(chan os.Signal, 1)
	//lint:ignore SA1016 i dont know, it just works lol
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-stop
	err = discord.Bot.Close()
	if err != nil {
		slog.Error("Cannot close discord bot, but we are going to close anyway")
	}

	slog.Info("Bye bye!")
}
