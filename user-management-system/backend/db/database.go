package db

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// InitDatabase initializes the database and applies migrations
func InitDatabase() (*sql.DB, error) {
	// SQLite connection string
	db, err := sql.Open("sqlite3", "./user_management.db")
	if err != nil {
		return nil, err
	}

	// Path to the migrations.sql file
	migrationFilePath := "../db/migrations.sql"

	// Read the SQL migration file
	migration, err := ioutil.ReadFile(migrationFilePath)
	if err != nil {
		return nil, err
	}

	// Execute the SQL commands from the migration file
	_, err = db.Exec(string(migration))
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully!")
	return db, nil
}
