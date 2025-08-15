package migration

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration(dsn string) {
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalf("failed to find migrations folder: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("migration running successfully")
}

func InsertInitialNationalities(db *sql.DB) {
	nationalities := []struct {
		Name string
		Code string
	}{
		{"Indonesia", "ID"},
		{"United States", "US"},
		{"Japan", "JP"},
	}

	var count int
	err := db.QueryRow(`select count(*) FROM nationality`).Scan(&count)
	if err != nil {
		log.Fatalf("failed to check nationality table: %v", err)
	}

	if count > 0 {
		log.Println("nationality table has been initiated")
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into nationality (nationality_name, nationality_code) VALUES ($1, $2)`)
	if err != nil {
		log.Fatalf("failed to prepare sql insertion: %v", err)
	}
	defer stmt.Close()

	for _, n := range nationalities {
		_, err := stmt.Exec(n.Name, n.Code)
		if err != nil {
			log.Fatalf("failed to execute nationality insertion: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("failed to commit: %v", err)
	}

	log.Println("insertion initial data for nationality is already success")
}
