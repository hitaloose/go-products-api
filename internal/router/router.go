package router

import "github.com/gin-gonic/gin"

func SetupRouter(server *gin.Engine) {
	v1 := server.Group("/api/v1")

	SetupProductRouter(v1)
}
