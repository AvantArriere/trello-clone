package main

import (
	"os"

	"github.com/jinzhu/gorm"
)

// MariaDB for database connections
type MariaDB struct {
	DB *gorm.DB
}

func (db *MariaDB) connect() (err error) {
	var connection *gorm.DB
	var dataString string

	if dataString := os.Getenv("DATABASE_STRING"); dataString == "" {
		dataString = "root:croissant123!/trello?charset=utf8,utf8mb4&parseTime=True&loc=Local"
	}

	if connection, err = gorm.Open("mysql", dataString); err == nil {
		db.DB = connection
	}
	return
}

func (db *MariaDB) close() error {
	return db.DB.Close()
}
