package main

import (
	"encoding/json"
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

func getMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["id"]
	found := false
	for _, item := range movies {
		if item.ID == id {
			found = true
			w.WriteHeader(http.StatusFound)
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, "enconding error", http.StatusInternalServerError)
				break
			}
		}
	}
	if !found {
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movies
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "error decoding r.body", http.StatusInternalServerError)
		return
	}
	movies = append(movies, movie)
	w.WriteHeader(http.StatusCreated)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movieFromBody Movies
	if err := json.NewDecoder(r.Body).Decode(&movieFromBody); err != nil {
		http.Error(w, "error decoding r.body", http.StatusInternalServerError)
		return
	}

	id := mux.Vars(r)["id"]
	for index, item := range movies {
		if item.ID == id {
			movies[index].Director = movieFromBody.Director
			movies[index].Isbn = movieFromBody.Isbn
			movies[index].Title = movieFromBody.Title
			w.WriteHeader(http.StatusAccepted)
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["id"]
	found := false
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func insertMovies() {
	movies = append(movies, Movies{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movies{ID: "2", Isbn: "448744", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
	movies = append(movies, Movies{ID: "3", Isbn: "448745", Title: "Movie Three", Director: &Director{FirstName: "Jane", LastName: "Doe"}})
	movies = append(movies, Movies{ID: "4", Isbn: "448746", Title: "Movie Four", Director: &Director{FirstName: "Mike", LastName: "Smith"}})
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"error": "method not allowed"}`))
}

func main() {
	insertMovies()
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	r.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)
	r.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
