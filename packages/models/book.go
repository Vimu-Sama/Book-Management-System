package models

import (
	"fmt"
	"log"

	"github.com/Vimu-Sama/Book-Management-System/packages/config"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var (
	db *gorm.DB
)

func Migrate() {
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() Book {
	if err := db.Create(book); err != nil {
		fmt.Printf("There was error while creating new Entry->  %v", err)
	}
	return *book
}

func GetBookById(userId int64) Book {
	var getBook Book
	if err := db.Where("ID=?", userId).Find(getBook); err != nil {
		log.Default().Printf("Error occured while Reading Id-> %v", err)
	}
	return getBook
}

func GetAllBooks() []Book {
	var getBookArray []Book
	if err := db.Find(getBookArray); err != nil {
		log.Default().Printf("Error occured while Reading Id-> %v", err)
	}
	return getBookArray
}

func DeleteBook(userId int64) Book {
	var deletedUser Book
	if err := db.Where("ID=?", userId).Delete(deletedUser); err != nil {
		log.Default().Printf("Error occured while deleting the User-> %v", err)
	}
	return deletedUser
}

func UpdateUser() {

}
