package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReplyCreated(context *gin.Context, response interface{}) {
	context.JSON(http.StatusCreated, response)
}

func ReplyOK(context *gin.Context, response interface{}) {
	context.JSON(http.StatusOK, response)
}
