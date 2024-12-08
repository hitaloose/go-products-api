package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hitaloose/go-products-api/internal/controllers"
)

func SetupProductRouter(router *gin.Engine) {
	productController := controllers.NewProductController()

	router.POST("/product", productController.Create)
}
