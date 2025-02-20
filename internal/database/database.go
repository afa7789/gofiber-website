package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	client *gorm.DB
}

// NewDatabase creates a new database connection struct
func NewDatabase() *Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	// creating DB URL
	URL := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USER,
		PASS,
		HOST,
		DBNAME)

	// connecting to DB
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		// fmt.Printf("Failed to connect to database! at %s", URL)
		return nil
	}
	print("Connected to database!")
	return &Database{
		client: db,
	}
}
