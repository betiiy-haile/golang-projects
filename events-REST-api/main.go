// package main

// import (
// 	"example.com/events-api/db"
// 	"example.com/events-api/routes"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	db.InitDB()
// 	server := gin.Default()  // sets up engine for us
// 	routes.RegisterRoutes(server)

// 	server.Run(":8080")  // localhost:8080
// }

package main

import (
	"example.com/events-api/db"
	"example.com/events-api/routes"

	"github.com/gin-gonic/gin"

	// Swagger-related imports
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "example.com/events-api/docs"
)

// @title Events API
// @version 1.0
// @description API for managing events and registrations.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db.InitDB()
	server := gin.Default()

	// Register all routes
	routes.RegisterRoutes(server)

	// Swagger endpoint
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	server.Run(":8080")
}

