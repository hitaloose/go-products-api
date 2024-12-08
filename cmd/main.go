package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hitaloose/go-products-api/internal/router"
)

func main() {
	server := gin.Default()

	router.SetupRouter(server)

	server.Run(":8080")
}
