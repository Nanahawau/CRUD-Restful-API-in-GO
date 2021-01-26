package main

import (
		// "fmt"
		"encoding/json"
	 	"log"
		"net/http"
		"math/rand"
		"strconv"
		"github.com/gorilla/mux"
	)

	//Book Struct 
	 type Book struct {
		ID   string  `JSON: "id"`
		Isbn   string  `JSON: "isbn"`
		Title   string  `JSON: "title"`
		Author   *Author  `JSON: "author"`
	 }

	 type Author struct {
		Firstname string `JSON: "firstName"`
		Lastname string `JSON: "lastName"`
		Email string `JSON: "email"`

	 }


	 // Init Book Variables as a slice book struct 
	 var books []Book

	 func getBooks (w http.ResponseWriter, r *http.Request)  {
		 w.Header().Set("Content-Type", "application/json")
		 json.NewEncoder(w).Encode(books) 
	 }

	 func getBook (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		for _, item := range books {
			 if item.ID == params["id"] {
				 json.NewEncoder(w).Encode(item)

				 return;
		}	
	}
	json.NewEncoder(w).Encode(&Book{})
	}
	func createBook (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book);
		book.ID = strconv.Itoa(rand.Intn(10000))
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		 
	}
	func updateBook (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range books {
			if item.ID == params["id"] {

				books = append(books[:index], books[index+1:]...)
				var book Book
				_ = json.NewDecoder(r.Body).Decode(&book);
				book.ID = params["id"]
				books = append(books, book)
				json.NewEncoder(w).Encode(book)
	   }	
			}
   }
	
	func deleteBook (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range books {
			if item.ID == params["id"] {

				books = append(books[:index], books[index+1:]...)
				break
	   }	
			}
	}

func main() {
	//init the router 
	r := mux.NewRouter()


	//Mock Data 
	books = append(books, Book {ID: "1", Isbn: "4323", 
	Title: "We are the gods", Author: &Author{Firstname: "Ali", Lastname: "Adeku", Email: "Ali@gmail.com"}})

	books = append(books, Book {ID: "2", Isbn: "4333", 
	Title: "We are the gods", Author: &Author{Firstname: "Nans", Lastname: "Adeku", Email: "Ali@gmail.com"}})
	books = append(books, Book {ID: "3", Isbn: "4343", 
	Title: "We are the gods", Author: &Author{Firstname: "Jumz", Lastname: "Adeku", Email: "Ali@gmail.com"}})

	//Route Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}