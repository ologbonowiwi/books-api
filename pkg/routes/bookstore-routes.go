package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/ologbonowiwi/bookstore-api/docs"
	"github.com/ologbonowiwi/bookstore-api/pkg/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	})
	router.GET("/books/", controllers.GetBook)
	router.POST("/books/", controllers.CreateBook)
	router.GET("/books/:id", controllers.GetBookById)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	router.PATCH("/books/:id", controllers.UpdateBookPartial)

	router.Run()
}
