package routes

import (
	"github.com/Vimu-Sama/Book-Management-System/packages/controllers"
	"github.com/gorilla/mux"
)

var BookManagementRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/", controllers.AddBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookRecord).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.RemoveBook).Methods("DELETE")
}
