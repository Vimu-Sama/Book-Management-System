package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Vimu-Sama/Book-Management-System/packages/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	appRouter := mux.NewRouter()
	routes.BookManagementRoutes(appRouter)
	http.Handle("/", appRouter)
	if err := http.ListenAndServe("LocalHost:"+os.Getenv("PORT"), appRouter); err != nil {
		log.Fatal(err)
	}
}
