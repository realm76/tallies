package app

import (
	"embed"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	encoder "tallies/internal"
	"tallies/internal/app/handlers"
)

//go:embed templates/*.gohtml
var templateFiles embed.FS

var templates = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

func addRoutes(logger *zap.SugaredLogger, mux *http.ServeMux) {
	enc := encoder.NewEncoder(templates)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok 1223"))
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/tallies", handlers.HandleGetTallies(logger, enc))
}
