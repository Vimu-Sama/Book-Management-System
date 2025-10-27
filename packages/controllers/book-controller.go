package controllers

import (
	"fmt"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Add Book Record Function called\n")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get Books Function called\n")
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get Book by Id Function\n")
}

func UpdateBookRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Update Book Record Function called\n")
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Delete Book Record Function called\n")
}
