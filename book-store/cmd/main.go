package main

import (
	"book-store/config"
	"book-store/models"
	"book-store/utils"
	"fmt"
)

func main() {
	fmt.Println("welcome to the bookstore")
	utils.CreateLettersList()
	fmt.Println(utils.GenerateID())
	config.ConnectDB()
	db := config.GetDB()
	db.AutoMigrate(&models.Book{})

	db.Create(&models.Book{
		Id:     utils.GenerateID(),
		Title:  "C Programming",
		Author: "Charls babege",
		Price:  525.23,
	})
}
