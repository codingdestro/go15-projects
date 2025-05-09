package config

import (
	"fms/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	fmt.Println("Connecting with database!")
	db, err := gorm.Open(sqlite.Open("fms.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect with database!")
		return
	}
	db.AutoMigrate(models.Users{})
	db.AutoMigrate(models.Folders{})
	db.AutoMigrate(models.Files{})
	DB = db
	fmt.Println("Connected with database!")
}
