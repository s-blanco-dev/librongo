package api

import (
	"librongo/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(bookHandler *handler.BookHandler, authorHandler *handler.AuthorHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/books", func(r chi.Router) {
		r.Get("/{id}", bookHandler.GetBookByID)
		r.Post("/add", bookHandler.CreateBook)
	})

	r.Route("/author", func(r chi.Router) {
		r.Get("/{id}", authorHandler.GetAuthorByID)
		r.Post("/add", authorHandler.CreateAuthor)
	})

	return r
}
