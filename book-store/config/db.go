package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	DB, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = DB
}

func GetDB() *gorm.DB {
	return db
}
