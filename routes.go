package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {

	a.App.Routes.Get("/", a.Handlers.Home)

	fileServer := http.FileServer(http.Dir("./public"))
	// http.StripPrefix handler will strip HTTP request's url prefix then send it to correspoding handler
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
