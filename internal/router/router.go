package router

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	SetupProductRouter(router)
}
