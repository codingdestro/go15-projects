package main

import (
	"book-store/config"
	"book-store/models"
	"book-store/routes"
	"book-store/utils"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("welcome to the bookstore")
	utils.CreateLettersList()
	config.ConnectDB()
	db := config.GetDB()
	db.AutoMigrate(&models.Book{})
	router := routes.InitRoutes()
	http.ListenAndServe(":8080", router)
}
