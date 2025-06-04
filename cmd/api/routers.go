package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routers() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/api/healthcheck", app.healthcheckHandler)
	router.Post("/api/books", app.createBookHandler)
	router.Get("/api/books/{id}", app.showBookHandler)

	return router
}
