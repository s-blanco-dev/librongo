package main

import (
	"librongo/internal/db"
	"log"
)

func main() {
	database, err := db.NewSQLite("db/books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	log.Println("Conexión a la BD marcha joya")
}
