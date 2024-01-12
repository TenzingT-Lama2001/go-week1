package database

import (
	"database/sql"
	"log"
	"time"
)

// NewDB creates a new database connection
func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Password!?@#$123@tcp(localhost:3306)/blogs")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

// CloseDB closes the database connection
func CloseDB(db *sql.DB) {
	db.Close()
}
