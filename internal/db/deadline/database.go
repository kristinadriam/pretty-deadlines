package db

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "deadlines.db"

type Database struct {
	db *sql.DB
}

func InitDb() (*Database, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return nil, err
	}
	defer db.Close()

	// create table with deadlines
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS deadlines (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		deadline TEXT
		);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
		return nil, err
	}

	return &Database{db: db}, nil
}

func (database *Database) Exec(query string, args ...any) (sql.Result, error) {
	return database.db.Exec(query, args)
}
