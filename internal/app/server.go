package app

import (
	"go.uber.org/zap"
	"net/http"
)

func NewServer(logger *zap.SugaredLogger) http.Handler {
	mux := http.NewServeMux()

	var handler http.Handler = mux

	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infow("request", "method", r.Method, "url", r.URL.String(), "remote", r.RemoteAddr)
		mux.ServeHTTP(w, r)
	})

	addRoutes(logger, mux)

	return handler
}
