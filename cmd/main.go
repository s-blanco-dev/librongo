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

	router := api.SetupRoutes(bookHandler, authorHandler)

	http.ListenAndServe(":8080", router)
}
