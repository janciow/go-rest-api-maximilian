package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("failed to connect database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(0)
	createTables()
}

func createTables() error {
	createEventTableSQL := `CREATE TABLE IF NOT EXISTS events (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"user_id" INTEGER,
		"name" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"location" TEXT NOT NULL,
		"date" DATETIME NOT NULL
	  );`

	statement, err := DB.Prepare(createEventTableSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		panic("Could not create events table")
	}
	return nil
}
