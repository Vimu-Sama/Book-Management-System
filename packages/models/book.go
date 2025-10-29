package models

import (
	"github.com/Vimu-Sama/Book-Management-System/packages/config"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Migrate() {
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}
