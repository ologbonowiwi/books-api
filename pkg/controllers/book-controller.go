package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ologbonowiwi/bookstore-api/pkg/models"
)

// @Summary Get books
// @Description get all saved books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Failure 404 {string} string "No books found"
// @Router /books [get]
func GetBook(ctx *gin.Context) {
	books := models.GetAllBooks()

	if len(books) == 0 {
		ctx.JSON(http.StatusNotFound, "No books found")
		return
	}

	ctx.JSON(http.StatusOK, books)
}

// @Summary Get book by id
// @Description get book by id
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 500 {string} string Error parsing the id
// @Failure 404 {string} string "Book with id not found"
// @Router /books/{id} [get]
func GetBookById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error parsing the id")
		return
	}

	book := models.GetBookById(id)

	if book.ID == 0 {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("Book with id %d not found", id))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// @Summary Create book
// @Description create book
// @Accept json
// @Produce json
// @Param book body models.Book true "Book"
// @Success 201 {object} models.Book
// @Failure 400 {string} string "Bad request"
// @Router /books [post]
func CreateBook(ctx *gin.Context) {
	book := &models.Book{}
	ctx.BindJSON(book)

	book = models.CreateBook(book)

	ctx.JSON(http.StatusCreated, book)
}

// @Summary Delete book
// @Description delete book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 204
// @Router /books/{id} [delete]
func DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error parsing the id")
		return
	}

	models.DeleteBook(id)

	ctx.JSON(http.StatusNoContent, nil)
}

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
func UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error parsing the id")
		return
	}

	book := &models.Book{}
	if err := ctx.BindJSON(book); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	book.ID = uint(id)
	book.UpdateBook()

	ctx.JSON(http.StatusOK, book)
}
