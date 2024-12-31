package router

import (
	"github.com/bojurgess/bard/internal/handler"
	"github.com/bojurgess/bard/internal/middleware"
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth", handler.Authorize)
	mux.HandleFunc("/callback", handler.Callback)
	mux.HandleFunc("/{id}/currently_playing", handler.CurrentlyPlaying)

	return middleware.ChainMiddleware(mux, middleware.LoggingMiddleware)
}
