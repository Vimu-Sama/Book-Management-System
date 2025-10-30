package main

import (
	"net/http"
	"os"

	"github.com/Vimu-Sama/Book-Management-System/packages/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err:= godotenv.Load(); err!=nil{
		panic(log) ;
	}
	appRouter := mux.NewRouter() ;
	routes.BookManagementRoutes(appRouter);
	http.Handle("/", appRouter);
	if err:= http.ListenAndServe("LocalHost:"+ os.Getenv("PORT"), appRouter)
}
