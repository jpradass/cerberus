package db

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	err      error
	userhome string  = os.Getenv("HOME")
	db       *sql.DB = nil
)

func Connect() error {
	db, err = sql.Open("sqlite", fmt.Sprintf("%s/cerberus.db", userhome))
	if err != nil {
		db = nil
		return fmt.Errorf("there was an error connecting to sqlite db: %w", err)
	}

	if err := db.Ping(); err != nil {
		db = nil
		return fmt.Errorf("failed when pinging the db: %w", err)
	}

	return nil
}
