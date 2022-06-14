package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	for i, book := range Books {
		if strconv.Itoa(int(Books[i].Id)) == id {
			w.Header().Set("Content-Type", "application/json")

			encoder := json.NewEncoder(w)
			encoder.Encode(&book)

			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	newBook := new(Book)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(newBook)

	Books = append(Books, *newBook)

	w.WriteHeader(http.StatusCreated)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i := range Books {
		if strconv.Itoa(int(Books[i].Id)) == id {
			Books[i] = Books[len(Books)-1]
			Books = Books[:len(Books)-1]

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(404)
}
