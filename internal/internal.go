package internal

import (
	"github.com/mauriciofsnts/exodia/internal/config"
)

func Bootstrap() {
	config.ConfigLogger()
	config.Init()
	// discord.InitDiscord()
}
