package main

import (
	"github.com/ologbonowiwi/bookstore-api/pkg/routes"
)

// @title The Bookstore API
// @description This is a sample bookstore API
// @version 1.0
// @termsOfService http://swagger.io/terms/

// @contact.name Swagger API Team
// @contact.url http://swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	routes.SetupRoutes()
}
