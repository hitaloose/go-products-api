package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hitaloose/go-products-api/internal/router"
)

// @title           Go Products API
// @version         1.0
// @description     A modern, lightweight API for product management built with Go and the Gin framework. Features clean architecture, mocked repositories, validation, and custom error handling for a seamless development experience.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Hitalo Loose
// @contact.url    https://github.com/hitaloose
// @contact.email  hitaloose@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	server := gin.Default()

	router.SetupRouter(server)
	router.SetupSwagger(server)

	server.Run(":8080")
}
