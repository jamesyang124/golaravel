package handler

import (
	"net/http"

	"github.com/jamesyang124/celeritas"
)

type Handlers struct {
	App *celeritas.Celeritas
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("err rendering:", err)
	}
}

func (h *Handlers) GoTemplatePage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "go-template", nil)
	if err != nil {
		h.App.ErrorLog.Println("err rendering:", err)
	}
}

func (h *Handlers) JetTemplatePage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "jet-template", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("err rendering:", err)
	}
}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	data := "foo"

	// store it to cookie by default session type
	h.App.Session.Put(r.Context(), "bar", data)

	err := h.App.Render.JetPage(w, r, "sessions", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("err rendering:", err)
	}
}
