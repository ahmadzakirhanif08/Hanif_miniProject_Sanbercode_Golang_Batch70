package configs

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" 
)

var DB *sql.DB

func InitDB() *sql.DB {
	config := LoadDBConfig()
	
	var err error
	DB, err = sql.Open("postgres", config.DSN())
	if err != nil {
		log.Fatalf("Error saat koneksi ke database: %v", err)
	}

	// 3. Tes Koneksi
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error saat ping database: %v", err)
	}

	log.Println("Koneksi database berhasil!")
	return DB
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Koneksi database ditutup.")
	}
}