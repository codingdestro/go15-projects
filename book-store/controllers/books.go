package controllers

import (
	"book-store/config"
	"book-store/models"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var books []models.Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}
