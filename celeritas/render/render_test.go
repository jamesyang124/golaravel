package render

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "Error rendering go template /home"},
	{"go_page_no_template", "go", "no-file", true, "open ./testdata/views/no-file.page.tmpl: no such file or directory"},
	{"jet_page", "jet", "home", false, "Error rendering jet template /home"},
	{"jet_page_no_template", "jet", "no-file", true, "template /no-file.jet could not be found"},
	{"unknown_template_engine", "unknown", "home", true, "no rendering engine specified"},
}

func TestRender_Page(t *testing.T) {

	for _, p := range pageData {
		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}

		// recorder let later action to inspect its response
		w := httptest.NewRecorder()

		testRenderer.Renderer = p.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, p.template, nil, nil)
		if p.errorExpected {
			assert.EqualError(t, err, p.errorMessage)
		} else if err != nil {
			t.Error("unexpected error", err)
		}
	}

}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering go template /home", err)
	}

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	assert.EqualError(t, err, "open ./testdata/views/no-file.page.tmpl: no such file or directory")

}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering jet template /home", err)
	}

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	assert.EqualError(t, err, "template /no-file.jet could not be found")
}
