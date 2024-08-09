package handlers

import (
	_ "embed"
	"go.uber.org/zap"
	"net/http"
	encoder "tallies/internal"
	"time"
)

func HandleGetTallies(logger *zap.SugaredLogger, encoder encoder.HtmlEncoder) http.HandlerFunc {
	if logger == nil {
		panic("nil logger")
	}

	if encoder == nil {
		panic("nil encoder")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		encoder.EncodeHTML(w, r, http.StatusOK, "tallies", map[string]interface{}{
			"tallies": []map[string]interface{}{
				{
					"Name":  "time " + time.Now().Format("2006-01-02 15:04:05"),
					"count": 1,
				},
			},
		})
	}
}
