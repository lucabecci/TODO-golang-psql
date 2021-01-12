package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//GetConnection is a function for the database connection.
func GetConnection() *sql.DB {

	URI, _ := os.LookupEnv("DB_URI")

	db, err := sql.Open("postgres", URI)

	if err != nil {
		log.Fatal(err)
	}

	return db

}
