package main

import (
	"librongo/internal/api"
	"librongo/internal/db"
	"librongo/internal/handler"
	"librongo/internal/repository"
	"librongo/internal/services"
	"log"
	"net/http"
)

func main() {
	database, err := db.NewSQLite("../db/books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	bookRepo := repository.NewBookRepository(database)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	authorRepo := repository.NewAuthorRepository(database)
	authorService := services.NewAuthorService(authorRepo)
	authorHandler := handler.NewAuthorHandler(authorService)

	topicRepo := repository.NewTopicRepository(database)
	topicService := services.NewTopicService(topicRepo)
	topicHandler := handler.NewTopicHandler(topicService)

	editorialRepo := repository.NewEditorialRepository(database)
	editorialService := services.NewEditorialService(editorialRepo)
	editorialHandler := handler.NewEditorialHandler(editorialService)

	router := api.SetupRoutes(bookHandler, authorHandler, topicHandler, editorialHandler)

	http.ListenAndServe(":8080", router)
}
