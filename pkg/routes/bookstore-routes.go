package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ologbonowiwi/bookstore-api/pkg/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	})
	router.GET("/books/", controllers.GetBook)
	router.POST("/books/", controllers.CreateBook)
	router.GET("/books/:id", controllers.GetBookById)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
