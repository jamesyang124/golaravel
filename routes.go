package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {

	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoTemplatePage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetTemplatePage)

	fileServer := http.FileServer(http.Dir("./public"))
	// http.StripPrefix handler will strip HTTP request's url prefix then send it to correspoding handler
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
