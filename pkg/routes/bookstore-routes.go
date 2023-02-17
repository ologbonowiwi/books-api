package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ologbonowiwi/bookstore-api/pkg/controllers"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	})
	engine.GET("/books/", controllers.GetBook)
	engine.POST("/books/", controllers.CreateBook)
	engine.GET("/books/:id", controllers.GetBookById)
	engine.PUT("/books/:id", controllers.UpdateBook)
	engine.DELETE("/books/:id", controllers.DeleteBook)
}
