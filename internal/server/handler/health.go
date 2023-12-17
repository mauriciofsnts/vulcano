package handler

import (
	"net/http"

	"github.com/mauriciofsnts/vulcano/internal/discord"
)

func Health(w http.ResponseWriter, r *http.Request) {
	botIsOK := discord.Bot.IsAlive()

	if botIsOK {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusServiceUnavailable)

}
