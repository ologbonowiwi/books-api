package models

import (
	"fmt"

	"github.com/ologbonowiwi/bookstore-api/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

func init() {
	config.Init()

	db = config.GetDB()

	db.AutoMigrate(&Book{})
}

func CreateBook(book *Book) *Book {
	db.Create(book)

	fmt.Println(book)

	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) *Book {
	var book Book
	db.Where("ID=?", id).Find(&book)
	return &book
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}

func (book *Book) UpdateBook() *Book {
	db.Save(book)
	return book
}
