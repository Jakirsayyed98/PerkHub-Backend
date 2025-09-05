package connection

import (
	"PerkHub/constants"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

func MakePotgressConn() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		constants.PostgresHost, constants.PostgresPort, constants.PostgresUsername, constants.PostgresPassword, constants.PostgresDatabase)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err // Return error instead of logging and exiting
	}

	// Check if the connection is valid
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)           // max 20 connections open at once
	db.SetMaxIdleConns(5)            // keep up to 5 idle
	db.SetConnMaxLifetime(time.Hour) // recycle connections after 1h

	// Run migrations
	if err := goose.Up(db, "migrations"); err != nil {
		return nil, err
	}

	log.Println("Migrations applied successfully!")
	return db, nil
}
