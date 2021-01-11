package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//GetConnection is a function for the database connection.
func GetConnection() *sql.DB {
	connStr := "postgres://mkozqujp:oqBgYGiLfaY0qwJo7w4TvJ14iyC7D2uP@tuffi.db.elephantsql.com:5432/mkozqujp"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db

}
