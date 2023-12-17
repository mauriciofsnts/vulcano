package server

import (
	"github.com/go-chi/chi"
	"github.com/mauriciofsnts/vulcano/internal/server/handler"
)

func route(r *chi.Mux) {
	r.Get("/health", handler.Health)
}
