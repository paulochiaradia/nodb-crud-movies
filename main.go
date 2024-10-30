package main

import (
	"fmt"
	"log"
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
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movies

func getMovies(w http.ResponseWriter, r *http.Request) {}

func getMovie(w http.ResponseWriter, r *http.Request) {}

func createMovie(w http.ResponseWriter, r *http.Request) {}

func updateMovie(w http.ResponseWriter, r *http.Request) {}

func deleteMovie(w http.ResponseWriter, r *http.Request) {}

func insertMovies() {
	movies = append(movies, Movies{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movies{ID: "2", Isbn: "448744", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
	movies = append(movies, Movies{ID: "3", Isbn: "448745", Title: "Movie Three", Director: &Director{FirstName: "Jane", LastName: "Doe"}})
	movies = append(movies, Movies{ID: "4", Isbn: "448746", Title: "Movie Four", Director: &Director{FirstName: "Mike", LastName: "Smith"}})
}

func main() {
	insertMovies()
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	r.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
