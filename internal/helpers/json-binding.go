package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/hitaloose/go-products-api/internal/exceptions"
)

func BindJSON(context *gin.Context, input interface{}) error {
	if err := context.ShouldBindJSON(&input); err != nil {
		return exceptions.NewBadRequestException("Invalid request", nil)
	}

	return nil
}

func BindQuery(context *gin.Context, input interface{}) error {
	if err := context.ShouldBindQuery(&input); err != nil {
		return exceptions.NewBadRequestException("Invalid request", nil)
	}

	return nil
}
