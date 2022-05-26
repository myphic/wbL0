package postgres

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewClient() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASENAME"),
	)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
