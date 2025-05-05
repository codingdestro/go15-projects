package routes

import (
	"book-store/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetABook).Methods("POST")
	return router
}
