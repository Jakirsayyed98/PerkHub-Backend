package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

func MakePotgressConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		return nil, err // Return error instead of logging and exiting
	}

	// Check if the connection is valid
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Run migrations
	if err := goose.Up(db, "migrations"); err != nil {
		return nil, err
	}

	log.Println("Migrations applied successfully!")
	return db, nil
}
