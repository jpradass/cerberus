package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"

	dbmodels "github.com/jpradass/cerberus/models/db"
)

var (
	err       error
	userhome  string  = os.Getenv("HOME")
	db        *sql.DB = nil
	querySQL          = `SELECT value, is_binary, is_path FROM cerberus_den WHERE key = ?;`
	insertSQL         = `INSERT INTO cerberus_den (key, value, is_binary, is_path) VALUES (?, ?, ?, ?);`
	initSQL           = `CREATE TABLE IF NOT EXISTS cerberus_den (
  key TEXT PRIMARY KEY, 
  value TEXT, 
  is_binary INTEGER NOT NULL CHECK (is_path IN (0, 1)),
  is_path INTEGER NOT NULL CHECK (is_path IN (0, 1))
);`
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

func SaveInDen(entry *dbmodels.Entry) error {
	if err := checkConn(); err != nil {
		return err
	}

	_, err := db.Exec(insertSQL, entry.Key, entry.Value, entry.IsBinary, entry.IsPath)
	if err != nil {
		return fmt.Errorf("error: cerberus didn't want to take the object into its den: %w", err)
	}

	return nil
}

func GetFromDen(key string) (*dbmodels.Entry, error) {
	if err := checkConn(); err != nil {
		return nil, err
	}

	rows, err := db.Query(querySQL, key)
	if err != nil {
		return nil, fmt.Errorf("error: cerberus didn't want to give the object away: %w", err)
	}

	defer rows.Close()

	entry := new(dbmodels.Entry)
	for rows.Next() {
		if err := rows.Scan(&entry.Value, &entry.IsBinary, &entry.IsPath); err != nil {
			return nil, fmt.Errorf("error: cerberus didn't want to give the object away: %w", err)
		}
	}

	// fmt.Printf("value: %s", value)
	return entry, nil
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

func boolToInt(b bool) int {
	return map[bool]int{true: 1, false: 0}[b]
}
