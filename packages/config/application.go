package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the environment variable file")
	}

	userName := os.Getenv("USERDB")
	password := os.Getenv("PASSDB")
	host := os.Getenv("HOSTDB")
	port := os.Getenv("PORTDB")
	dbName := os.Getenv("NAMEDB")

	// ✅ Properly formatted MySQL DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		userName, password, host, port, dbName)

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = d
	log.Println("✅ Connected to MySQL successfully!")
}

func GetDB() *gorm.DB {
	return db
}
