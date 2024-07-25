package internal

import (
	"github.com/mauriciofsnts/exodia/internal/config"
	"github.com/mauriciofsnts/exodia/internal/discord"
)

func Bootstrap() {
	config.ConfigLogger()
	config.Init()
	discord.InitDiscord()
}
