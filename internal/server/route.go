package server

import (
	"github.com/go-chi/chi"
	"github.com/mauriciofsnts/bot/internal/server/handler"
)

func route(r *chi.Mux) {
	r.Get("/health", handler.Health)
}
