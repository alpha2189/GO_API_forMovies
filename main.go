package main

import(

"fmt"
"log"
"encoding/json"
"math/rand"
"net/http"
"strconv"
"github.com/gorilla/mux"

)


////Creating Movie Struct

type Movie struct{
	ID string 'json: "id"'
	Isbn string 'json:"isbn"'
	Title string 'json:"title"'
	Director *Director 'json:"director"'
}

////Creating Director Struct
type Director{
	Firstname string 'json:"firstname"'
	Lastname string 'json:"lastname"'
}

var movies []Movie

////function getMovies
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

////function deleteMovie
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
		}
	}

}

////function getMovie
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content - Type", "application/json")
	params := mux.Vars(r) 
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(movies)
}


////function createMovie
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content - Type", "application/json")
	var movie Movie_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, Movie)
	json.NewEncoder(w).Encode(movie)

}


////function updateMovie
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content - Type", "application/json")	
	params :=mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder().Encode(movie)
			return 
		}
	}
}


////Func main of the programm
func main(){
r := mux.NewRouter()

movies = append(movies, Movie{ID:"1", ISBN:"121212", Title:"First Movie", Director: &Director{Firstname:"Steve", Lastname: "Pap"}})
movies = append(movies, Movie{ID:"2", ISBN:"111111", Title:"Second Movie", Director: &Director{Firstname:"Elon", Lastname: "Tsi"}})

r.HandleFunc("/movies", getMovies).Method("GET") 
r.HandleFunc("/movies/{id}", getMovie).Method("GET")
r.HandleFunc("/movies", createMovie).Method("POST")
r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")

fmt.Printf("starting server at port 8080\n")
log.Fatal(http.ListenAndServe(":8080", r))

}
