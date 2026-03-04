package http

import (
	"net/http"

	"github.com/ArtoIi/Blogging-Platform-API/internal/interfaces/http/handlers"
)

func SetupRoutes(handler *handlers.PostHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /posts", handler.Create)
	mux.HandleFunc("GET /posts", handler.GetAll)
	mux.HandleFunc("GET /posts/{id}", handler.GetByID)
	mux.HandleFunc("PUT /posts/{id}", handler.Update)
	mux.HandleFunc("DELETE /posts/{id}", handler.Delete)

	return mux
}
