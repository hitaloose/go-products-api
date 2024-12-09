package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hitaloose/go-products-api/cmd/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupSwagger(server *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
