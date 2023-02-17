package models

import (
	"github.com/ologbonowiwi/bookstore-api/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	ID          uint
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type bookModel struct {
	gorm.Model
	Book
}

var db *gorm.DB

func init() {
	config.Init()

	db = config.GetDB()

	db.AutoMigrate(&bookModel{})
}

func (book *Book) parseToModel() *bookModel {
	model := bookModel{}

	model.Author = book.Author
	model.Name = book.Name
	model.Publication = book.Publication

	return &model
}

func CreateBook(book *Book) *Book {
	db.Create(book.parseToModel())
	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) *Book {
	var book Book
	db.Where("ID=?", id).Find(&book)
	return &book
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book.parseToModel())
	return book
}

func (book *Book) UpdateBook() *Book {
	db.Model(book.parseToModel()).Updates(book)

	return book
}
