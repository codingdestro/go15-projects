package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Not a vaild form", http.StatusNotAcceptable)
		return
	}

	username := r.FormValue("username")
	rollno := r.FormValue("rollno")
	age := r.FormValue("age")

	fmt.Printf("username = %s\n", username)
	fmt.Printf("username = %s\n", age)
	fmt.Printf("username = %s\n", rollno)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" {
			http.Error(w, "Not a valid route!", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "this is you hello function")
	})

	http.HandleFunc("/form", handleForm)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
