package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ologbonowiwi/bookstore-api/pkg/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
