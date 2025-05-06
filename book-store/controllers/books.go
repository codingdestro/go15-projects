package controllers

import (
	"book-store/config"
	"book-store/models"
	"book-store/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var books []models.Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "no book got")
		return
	}
	var book models.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		fmt.Fprintln(w, "no book got")
		return
	}
	book.Id = utils.GenerateID()
	fmt.Println(book)
	result := db.Create(book)
	if result.Error != nil {
		fmt.Fprint(w, "Failed to add new book!")
		return
	}
	fmt.Fprint(w, "your book saved!")

}

func GetABook(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	params := mux.Vars(r)
	bookId := params["id"]
	var book models.Book
	res := db.Where("id = ?", bookId).First(&book)
	if res.Error != nil {
		fmt.Fprintf(w, "failed to fetch you book id = %s", bookId)
		return
	}
	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Fprintf(w, "failed to fectch your book")
		return
	}
	json.NewEncoder(w).Encode(jsonData)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	db := config.GetDB()

	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "No data found!")
		return
	}

	var newBook models.Book
	err = json.Unmarshal(buffer, &newBook)
	if err != nil {
		fmt.Fprint(w, "No data found!")
		return
	}

	var book models.Book
	db.Where("id = ?", id).First(&book)

	book.Author = newBook.Author
	book.Title = newBook.Title
	book.Price = newBook.Price
	res := db.Save(book)
	if res.Error != nil {
		fmt.Fprint(w, "Failed to update book!")
		return
	}

	fmt.Fprint(w, "Book has been updated!")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	db := config.GetDB()
	db.Where("id = ?", id).Delete(&models.Book{})
	fmt.Fprintf(w, "Book has been deleted successfully")
}
