package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"io"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
	"strconv"
	_ "strconv"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

// movie structure
type Movie struct {
	Id       string    `"json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director structure
type Director struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

// array of the movie structure
var movies []Movie

var _id = 0

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var mov Movie
	json.NewDecoder(r.Body).Decode(&mov)
	mov.Id = strconv.Itoa(_id)
	_id++
	movies = append(movies, mov)

	// print the body
	bodybyte, _ := io.ReadAll(r.Body)
	strbody := string(bodybyte)
	fmt.Println("body: " + strbody)

	// send response
	json.NewEncoder(w).Encode(mov)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for idx, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var bodyMov Movie
	json.NewDecoder(r.Body).Decode(&bodyMov)

	for idx, mov := range movies {
		if mov.Id == params["id"] {
			bodyMov.Id = params["id"]
			movies[idx] = bodyMov
			json.NewEncoder(w).Encode(bodyMov)
			return
		}
	}
}

func main() {
	// new instance of the mux router
	r := mux.NewRouter()

	// Route for getting all the movies
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// Route for getting the movie with specific id
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// Route for adding the new movies
	r.HandleFunc("/movies", addMovie).Methods("POST")
	// Route for deleting the movie with specific id
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	// Route for updating the movie with specific id
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	// Start the Server
	fmt.Println("Starting the server at port :3000...\n")
	log.Fatal(http.ListenAndServe(":3000", r))
}
