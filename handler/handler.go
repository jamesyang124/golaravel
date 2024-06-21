package handler

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
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
	h.App.Session.Put(r.Context(), "foo", data)

	value := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", value)

	err := h.App.Render.JetPage(w, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("err rendering:", err)
	}
}
