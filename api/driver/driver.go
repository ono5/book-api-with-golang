package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB - connect to postgres container
func ConnectDB() *sql.DB {
	var err error
	pgURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	return db
}
