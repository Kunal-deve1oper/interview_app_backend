package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func InitDB() {
	// Get the PostgreSQL connection string from environment variables
	dataSourceName := os.Getenv("DB_CONNECTION_STRING")

	var err error
	// Open a PostgreSQL connection
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Could not open db connection: %v", err)
	}

	// Verify the connection with a ping
	if err = DB.Ping(); err != nil {
		log.Fatalf("Could not ping db: %v", err)
	}

	log.Println("Connected to PostgreSQL database")
}
