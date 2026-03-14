package api

import (
	"librongo/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(bookHandler *handler.BookHandler, authorHandler *handler.AuthorHandler, topicHandler *handler.TopicHandler, editorialHandler *handler.EditorialHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/books", func(r chi.Router) {
		r.Get("/{id}", bookHandler.GetBookByID)
		r.Post("/add", bookHandler.CreateBook)
		r.Get("/all", bookHandler.GetAllBooks)
	})

	r.Route("/author", func(r chi.Router) {
		r.Get("/", authorHandler.GetAllAuthors)
		r.Get("/{id}", authorHandler.GetAuthorByID)
		r.Post("/add", authorHandler.CreateAuthor)
	})

	r.Route("/topic", func(r chi.Router) {
		r.Get("/", topicHandler.GetAllTopics)
		r.Post("/", topicHandler.CreateTopic)
	})

	r.Route("/editorial", func(r chi.Router) {
		r.Get("/", editorialHandler.GetAllEditorials)
		r.Post("/", editorialHandler.CreateEditorial)
	})

	return r
}
