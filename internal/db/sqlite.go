package db

import (
	"database/sql"
	"fmt"
)

func NewSQLite(path string) (*sql.DB, error) {
	/* Data Source Name (DSN), segun documentacion: "When creating a new SQLite database or connection to an existing one, with the file name additional options can be given." */
	dsn := fmt.Sprintf("%s?_foreign_keys=on", path) // activo claves foráenas

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	// por las dudas
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
