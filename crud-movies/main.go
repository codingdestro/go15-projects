package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movies struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movies = []Movies{
	{
		ID:    "1",
		Isbn:  "978-3-16-148410-0",
		Title: "The Shawshank Redemption",
		Director: &Director{
			Firstname: "Frank",
			Lastname:  "Darabont",
		},
	},
	{
		ID:    "2",
		Isbn:  "978-1-23-456789-0",
		Title: "The Godfather",
		Director: &Director{
			Firstname: "Francis",
			Lastname:  "Coppola",
		},
	},
	{
		ID:    "3",
		Isbn:  "978-0-12-345678-9",
		Title: "Pulp Fiction",
		Director: &Director{
			Firstname: "Quentin",
			Lastname:  "Tarantino",
		},
	},
	{
		ID:    "4",
		Isbn:  "978-9-87-654321-0",
		Title: "The Dark Knight",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:    "5",
		Isbn:  "978-4-56-789012-3",
		Title: "Inception",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")

	fmt.Println(("welcome to movies server!"))
	http.ListenAndServe(":8080", router)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, v := range movies {
		if v.ID == id {
			movie := v
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	fmt.Fprint(w, "sorry :( movie not found")
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movies
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		fmt.Fprint(w, "failed to parse json data!")
	}

	movies = append(movies, movie)

	fmt.Fprint(w, "new movie added to server!")
}
