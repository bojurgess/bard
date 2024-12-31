package database

import (
	"fmt"
	"github.com/bojurgess/bard/internal/config"
	"github.com/jmoiron/sqlx"
	"log"
	_ "modernc.org/sqlite"
)

var db *sqlx.DB

func Initialize() error {
	var err error
	db, err = sqlx.Open("sqlite", config.AppConfig.DatabaseUrl)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Connected to database")
	return nil
}
