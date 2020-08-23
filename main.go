package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//!declares the properties & the json field brings the field from the Json

//Book Model
type Book struct {
	ID     string  `json:"id"`     //Json
	Isbn   string  `json:"isbn"`   //Json
	Title  string  `json:"title"`  //Json
	Author *Author `json:"author"` //Json
}

//Author Model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as slice Book struct
var books []Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //! GET PARAMS

	//Looping through the list to find the book with that id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100)) //!MOCK ID
	books = append(books, book)

	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//!Initialize the Router
	r := mux.NewRouter()

	//TODO: implement database
	//!Mock Data
	books = append(
		books,
		Book{
			ID:    "1",
			Isbn:  "123123",
			Title: "Game of thrones",
			Author: &Author{
				Firstname: "George",
				Lastname:  "martin",
			},
		},
	)
	books = append(
		books,
		Book{
			ID:    "2",
			Isbn:  "19923",
			Title: "Devlok",
			Author: &Author{
				Firstname: "Dev",
				Lastname:  "Dutt",
			},
		},
	)

	//!Route handlers which will establish our endpoints for the api
	//!Methods defines the type of methods we want to use
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//!To run the server we use HTTP package
	log.Fatal(http.ListenAndServe(":8000", r))

}
