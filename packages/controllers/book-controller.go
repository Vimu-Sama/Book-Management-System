package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Vimu-Sama/Book-Management-System/packages/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Add Book Record Function called\n")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get Books Function called\n")
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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
