package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var (
	err       error
	userhome  string  = os.Getenv("HOME")
	db        *sql.DB = nil
	querySQL          = `SELECT value FROM cerberus_den WHERE key = ?;`
	insertSQL         = `INSERT INTO cerberus_den (key, value) VALUES (?, ?);`
	initSQL           = `CREATE TABLE IF NOT EXISTS cerberus_den (key TEXT PRIMARY KEY, value TEXT);`
)

func Connect() error {
	db, err = sql.Open("sqlite", fmt.Sprintf("%s/cerberus.den", userhome))
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

func CreateDen() error {
	if err := checkConn(); err != nil {
		return err
	}

	_, err := db.Exec(initSQL)
	if err != nil {
		return fmt.Errorf("error: couldn't walk into cerberus den: %w", err)
	}

	return nil
}

func SaveInDen(key, value string) error {
	if err := checkConn(); err != nil {
		return err
	}

	_, err := db.Exec(insertSQL, key, value)
	if err != nil {
		return fmt.Errorf("error: cerberus didn't want to take the object into its den: %w", err)
	}

	return nil
}

func GetFromDen(key string) (string, error) {
	if err := checkConn(); err != nil {
		return "", err
	}

	rows, err := db.Query(querySQL, key)
	if err != nil {
		return "", fmt.Errorf("error: cerberus didn't want to give the object away: %w", err)
	}

	defer rows.Close()

	var value string
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			return "", fmt.Errorf("error: cerberus didn't want to give the object away: %w", err)
		}
	}

	// fmt.Printf("value: %s", value)
	return value, nil
}

func CheckDenExistence() bool {
	cerb_path := fmt.Sprintf("%s/cerberus.den", userhome)

	_, err := os.Stat(cerb_path)
	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}

func GetCerberusDenPath() string {
	return fmt.Sprintf("%s/cerberus.den", userhome)
}

func checkConn() error {
	if db == nil {
		if err := Connect(); err != nil {
			return err
		}
	}
	return nil
}
