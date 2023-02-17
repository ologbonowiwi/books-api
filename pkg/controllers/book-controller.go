package controllers

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get books
// @Description get all saved books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Failure 404 {string} string "No books found"
// @Router /books [get]
func GetBook(ctx *gin.Context) {}

// @Summary Get book by id
// @Description get book by id
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {string} string "Book with id not found"
// @Router /books/{id} [get]
func GetBookById(ctx *gin.Context) {}

// @Summary Create book
// @Description create book
// @Accept json
// @Produce json
// @Param book body models.Book true "Book"
// @Success 201 {object} models.Book
// @Failure 400 {string} string "Bad request"
// @Router /books [post]
func CreateBook(ctx *gin.Context) {}

// @Summary Delete book
// @Description delete book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 204
// @Router /books/{id} [delete]
func DeleteBook(ctx *gin.Context) {}

// @Summary Update book
// @Description update book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Book with id not found"
// @Router /books/{id} [put]
func UpdateBook(ctx *gin.Context) {}
