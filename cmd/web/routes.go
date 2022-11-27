package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mcgigglepop/golang-bookings-website/pkg/config"
	"github.com/mcgigglepop/golang-bookings-website/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// handle static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
