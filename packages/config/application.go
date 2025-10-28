package config

//adding the DB connection here
import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the enviroment variable file")
	}
	userName := os.Getenv("USERDB")
	password := os.Getenv("PASSDB")
	d, err := gorm.Open("mysql", userName+":"+password+"?charset=UTF-8")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
