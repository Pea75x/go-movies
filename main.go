package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
		ID string `json:"id"`
		Isbn string `json:"isbn"`
		Title string `json:"title"`
		Director *Director `json:"director"`
}

type Director struct {
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
		w.Header().set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)
		// Encodes movies as JSON and writes it to w
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
		w.Header().set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range movies {
				if item.ID == params["ID"]{
						movies = append(movies[:index], movies[index + 1:]...)
						// delete movie in array - by replacing the movie with all of the movies that come after it
						break
				}
		}
		json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
		w.Header().set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, item := range movies {
				if item.ID == params["ID"] {
						json.NewEncoder(w).Encode(item)
						return
				}
		}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
		w.Header().set("Content-Type", "application/json")
		var movie Movie
		_ = json.NewDecoder(r.body).Decode(&movie)
		// turns JSON data into Go struct and saves it where the pointer for movie is
		movie.ID = strconv.Itoa(rand.Intn(100000000))
		movies.append(movies, movie)
		json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
		w.Header().set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range movies {
				if item.ID == params["ID"]{
						// Delete movie
						movies = append(movies[:index], movies[index + 1:]...)
						// Add new movie
						var movie Movie
						_ = json.NewDecoder(r.body).Decode(&movie)
						movie.ID = strconv.Itoa(rand.Intn(100000000))
						movies.append(movies, movie)
						json.NewEncoder(w).Encode(movie)
				}
		}
}

func main() {
		r := mux.NewRouter()

		movies = append(movies, Movie{ID: "1", Isbn: "43288", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
		movies = append(movies, Movie{ID: "2", Isbn: "86773", Title: "Movie 2", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

		r.HandleFunc("/movies", getMovies).methods("GET")
		r.HandleFunc("/movies/{id}", getMovie).methods("GET")
		r.HandleFunc("/movies", createMovie).methods("POST")
		r.HandleFunc("/movies/{id}", updateMovie).methods("PUT")
		r.HandleFUnc("/movies/{id}", deleteMovie).methods("DELETE")

		fmt.PrintF("Starting server at port 8000\n")
		log.Fatal(http.ListenAndServe(":8000", r))
}