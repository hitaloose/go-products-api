package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinExceptionHandler(context *gin.Context, err error) {
	if httpErr, ok := err.(*HttpException); ok {
		context.JSON(httpErr.StatusCode, gin.H{"message": httpErr.Message, "details": httpErr.Details})
		context.Abort()

		return
	}

	context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "details": err.Error()})
	context.Abort()
}
