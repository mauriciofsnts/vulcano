package internal

import (
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/discord"
)

func Bootstrap() {
	config.ConfigLogger()
	config.Init()
	discord.Init()
}
