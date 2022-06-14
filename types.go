package main

type Book struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books []Book = make([]Book, 0)
