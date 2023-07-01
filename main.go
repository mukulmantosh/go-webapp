package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       int       `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: 1, Isbn: "232323",
		Title:    "Sample Movie",
		Director: &Director{Firstname: "Alex", Lastname: "Jones"}})

	movies = append(movies, Movie{ID: 2, Isbn: "33332323",
		Title:    "Sample Movie2",
		Director: &Director{Firstname: "James", Lastname: "Bond"}})

	movies = append(movies, Movie{ID: 3, Isbn: "245332323",
		Title:    "Sample Movie3",
		Director: &Director{Firstname: "Dwayne", Lastname: "Johnson"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	Id, _ := strconv.Atoi(params["id"])
	for index, item := range movies {
		if item.ID == Id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	Id, _ := strconv.Atoi(params["id"])
	for _, item := range movies {
		if item.ID == Id {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	writer.WriteHeader(http.StatusNotFound)
	json.NewEncoder(writer).Encode(map[string]string{"message": "Not Found"})

}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	Id, _ := strconv.Atoi(params["id"])
	for index, item := range movies {
		if item.ID == Id {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = rand.Intn(100000)
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movies)
		}
	}

}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = rand.Intn(100000)
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movies)
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)

}
