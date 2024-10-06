package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
