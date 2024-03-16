package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	return db, err
}
