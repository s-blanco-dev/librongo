package api

import (
	"librongo/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(bookHandler *handler.BookHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/books", func(r chi.Router) {
		r.Get("/{id}", bookHandler.GetBookByID)
	})

	return r
}
