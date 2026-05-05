package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := "user=postgres password=Admin dbname=taskdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
