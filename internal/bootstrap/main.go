package bootstrap

import (
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord"
	"github.com/pauloo27/logger"
)

func Start() {
	logger.HandleFatal(config.LoadConfig(), "Failed to load config, check if config.yaml exists")
	logger.HandleFatal(discord.Start(), "Failed to start discord")
}
