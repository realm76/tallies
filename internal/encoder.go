package encoder

import (
	"html/template"
	"net/http"
)

type Encoder struct {
	templates *template.Template
}

type HtmlEncoder interface {
	EncodeHTML(w http.ResponseWriter, r *http.Request, status int, name string, data any)
}

func NewEncoder(templates *template.Template) *Encoder {
	if templates == nil {
		panic("nil templates")
	}

	return &Encoder{
		templates: templates,
	}
}

func (e *Encoder) EncodeHTML(w http.ResponseWriter, r *http.Request, status int, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := e.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
