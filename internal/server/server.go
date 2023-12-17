package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/pauloo27/logger"
)

func run() error {
	logger.Info("Starting HTTP server...")

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
		Addr:         fmt.Sprintf(":%s", config.Vulcano.Port),
	}

	logger.Infof("HTTP server started on port %s", config.Vulcano.Port)

	return server.ListenAndServe()
}

func StartHttpServer() {
	logger.HandleFatal(run(), "Failed to start HTTP server")
}
