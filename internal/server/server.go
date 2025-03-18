package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mauriciofsnts/bot/internal/config"
)

func run(port int16) error {
	slog.Info("Starting HTTP server...")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	route(r)

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", port),
	}

	slog.Info("HTTP server started on port ", "port", port)

	return server.ListenAndServe()
}

func StartHttpServer(cfg config.Config) {
	err := run(cfg.Server.Port)

	if err != nil {
		slog.Error("Failed to start HTTP server: ", "error", err)
		panic(err)
	}
}
