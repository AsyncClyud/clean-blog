package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDataBase(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
