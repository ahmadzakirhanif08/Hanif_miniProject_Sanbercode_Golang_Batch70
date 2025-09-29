package main

import (
    "log"
    "miniProject/config"
    "miniProject/routes" // <-- IMPORT PACKAGE ROUTES ANDA
    "github.com/joho/godotenv"
    "github.com/rubenv/sql-migrate"
    // "github.com/gin-gonic/gin" // <-- Tambahkan import gin
)

func main() {
	//0. check env file
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, assuming production environment.")
    }

    // 1. Inisialisasi Koneksi Database
    db, err := config.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 

    // 2. Jalankan Migrasi Database
    migrations := &migrate.FileMigrationSource{
        Dir: "migrations", 
    }

    // Jalankan UP (membuat tabel)
    n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }
    log.Printf("Applied %d migrations!", n)
    
    // 3. Setup dan Jalankan Router Gin
    router := routes.SetupRouter() // <-- Panggil fungsi SetupRouter dari package routes

    log.Println("Server starting on :8787")
	router.SetTrustedProxies([]string{"127.0.0.1"})
    // Pastikan Anda menangani error di Run jika ada
    if err := router.Run(":8787"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
	
}