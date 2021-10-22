package main

import (
	"net/http"

	"github.com/bartvanbenthem/gofound-restful/internal/config"
	"github.com/bartvanbenthem/gofound-restful/internal/handlers"
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

	// public routes
	router.Group(func(router chi.Router) {
		router.Get("/status", handlers.Repo.Status)
		router.Get("/v1/software", handlers.Repo.GetAllSoftware)
		router.Get("/v1/software/{id}", handlers.Repo.GetSoftwareByID)
		router.Get("/v1/categories", handlers.Repo.GetAllCategories)
		router.Get("/v1/categories/{category_id}", handlers.Repo.GetAllSoftwareByCategory)

		router.Post("/v1/login", handlers.Repo.Login)
	})

	// protected routes
	router.Group(func(router chi.Router) {
		router.Use(TokenVerify)
		router.Get("/v1/admin/deletesoftware/{id}", handlers.Repo.DeleteSoftware)
		router.Post("/v1/admin/editSoftware", handlers.Repo.EditSoftware)
		router.Post("/v1/admin/signup", handlers.Repo.Signup)
	})

	return router
}
