package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Vimu-Sama/Book-Management-System/packages/models"
	"github.com/Vimu-Sama/Book-Management-System/packages/utils"
	"github.com/gorilla/mux"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Add Book Record Function called\n")
	var newBook models.Book
	utils.ParseBody(r, &newBook)
	bookDetails := newBook.CreateBook()
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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
	vars := mux.Vars(r)
	extractedId := vars["BookId"]
	bookId, _ := strconv.ParseInt(extractedId, 0, 0)
	bookDetails, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Update Book Record Function called\n")
	var bookUpdatedData models.Book
	data := mux.Vars(r)
	idString := data["BookId"]
	bookId, _ := strconv.ParseInt(idString, 0, 0)
	// _ = models.DeleteBook(bookId)
	bookDetails, db := models.GetBookById(bookId)
	utils.ParseBody(r, bookUpdatedData)
	if bookUpdatedData.Name != "" {
		bookDetails.Name = bookUpdatedData.Name
	}
	if bookUpdatedData.Author != "" {
		bookDetails.Author = bookUpdatedData.Author
	}
	if bookUpdatedData.Publication != "" {
		bookDetails.Publication = bookUpdatedData.Publication
	}
	db.Save(&bookDetails)
	resMarshalled, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resMarshalled)
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, _ := strconv.ParseInt(vars["BookId"], 0, 0)
	deletedBook := models.DeleteBook(bookId)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
