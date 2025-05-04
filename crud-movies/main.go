package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(("welcome to movies server!"))
	router := mux.NewRouter()
	router.HandleFunc("/", getMovies).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func getMovies(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "get you movies")

}
