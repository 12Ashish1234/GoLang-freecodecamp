package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct of type Movie
type Movie struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	// Here * denotes a pointer. It is pointing to the Director struct that has been created. And its is now associated with the Movie struct
	Director *Director `json:"director"`
}

// Struct of type Director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// slice of type movies
var movies []Movie

/*
 * Get all movies in the slice
 */
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

/*
 * Delete a movie based on given ID
 */
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Here basically the movie which is there in movies[index], will be appended by
			// the movies[index+1] and the subsequent movies which are present after that. so
			// This will help deleting the movie which in at the index. So now the movie in index will cease to exist.
			break
		}
	}
	// This line will display the movies list
	json.NewEncoder(w).Encode(movies)
}

/*
 * Get a movie based on given ID
 */
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Here in this for loop we will not be using the indexes. Hence we dont need to use that variable. so we have left it blank
	for _, item := range movies {
		if item.ID == params["id"] {
			/// this will display the movie based on the id we submitted
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

/*
 * Create a movie
 */
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Defining a variable movie of type Movie
	var movie Movie

	// blank identifier. Here r.Body is basically the complete json body which will be sent here through the request body of the json
	// So we need to decode it.
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Now in order to create a new movie entry, we first have to generate a random movie ID. as ID is a field in movie struct.
	// So we are using rand.Intn which generates a random postive integer within the defined range. example 1000000 as shown below.
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	// the movie which was POSTED will be appended to the movies slice using append
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

/*
 * Update a movie based on given ID.
 * Here we will be first deleting the movie with the given ID and next will be adding the updated movie details.
 * This is not the right way. but we will be able to do proper coding later using Databases.
 */
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//delete the movie now
	for index, item := range movies {
		if item.ID == params["id"] {
			// This below line is going to delete the movie
			movies = append(movies[:index], movies[index+1:]...)

			// now this is going to create a movie in the same ID. The logic is same as createMovie()
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			//This below line is the only difference compared to createMovie() logic, because we want to use the same ID.
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)

		}
	}
}

func main() {
	r := mux.NewRouter() //This returns a new mux router instance using gorilla mux library

	// Hardcoding two movies so that we could test if the server is working fine
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "955976", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Wolfe"}})

	// HandleFunc is to handle the API request of particular types and route them to the mentioned handler functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Startingserver at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
