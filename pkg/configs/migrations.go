package configs

import (
	"database/sql"
	"log"
	"os"
	"github.com/rubenv/sql-migrate"
)


func RunMigrations(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations", 
	}

	dialect := "postgres"

	mode := os.Getenv("MIGRATION_MODE")

	if mode == "DOWN" {
		n, err := migrate.Exec(db, dialect, migrations, migrate.Down)
		if err != nil {
			log.Fatalf("Gagal saat melakukan rollback migrasi: %v", err)
		}
		log.Printf("Berhasil melakukan rollback %d migrasi.", n)
		return
	}

	n, err := migrate.Exec(db, dialect, migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Gagal saat menjalankan migrasi UP: %v", err)
	}

	if n == 0 {
		log.Println("Database sudah terbarukan (0 migrasi diterapkan).")
	} else {
		log.Printf("Berhasil menerapkan %d migrasi!", n)
	}
}