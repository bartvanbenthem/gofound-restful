package main

import (
	"net/http"

	"github.com/bartvanbenthem/gofound-restfull/internal/config"
	"github.com/bartvanbenthem/gofound-restfull/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.StripSlashes)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(enableCORS)

	router.Get("/status", handlers.Repo.Home)
	router.Get("/v1/software", handlers.Repo.GetAllSoftware)
	router.Get("/v1/software/{id}", handlers.Repo.GetSoftwareByID)

	return router
}
