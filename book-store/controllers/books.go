package controllers

import (
	"book-store/config"
	"book-store/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var books []models.Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func GetABook(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	params := mux.Vars(r)
	bookId := params["id"]
	var book models.Book
	db.Where("id = ?", bookId).First(&book)
	json.NewEncoder(w).Encode(book)
}
