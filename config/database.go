package config

import (
	"database/sql"
	"fmt"
	"os"
	
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	var database string
	if os.Getenv("DATABASE_URL") != "" {
		database = os.Getenv("DATABASE_URL")
	} else {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		database = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbHost, dbUser, dbPassword, dbName, dbPort)
	}

	var err error
	DB, err = sql.Open("postgres", database)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Ping untuk memverifikasi koneksi
	if err = DB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Successfully connected to the database!")
	return DB, nil
}
