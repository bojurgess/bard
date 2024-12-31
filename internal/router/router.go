package router

import (
	"github.com/bojurgess/bard/internal/handler"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth", handler.Authorize)

	return mux
}
