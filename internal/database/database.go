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
		panic("Failed to connect to database!")
	}

	return &Database{
		client: db,
	}
}
