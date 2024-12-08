package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hitaloose/go-products-api/internal/dtos"
	"github.com/hitaloose/go-products-api/internal/exceptions"
	"github.com/hitaloose/go-products-api/internal/helpers"
	"github.com/hitaloose/go-products-api/internal/services"
)

type ProductController struct {
	createProductService services.CreateProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		createProductService: *services.NewCreateProductService(),
	}
}

func (controller ProductController) Create(context *gin.Context) {
	var body dtos.CreateProductDto

	if err := helpers.BindJSON(context, body); err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	if err := helpers.ValidateStruct(body); err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	product, err := controller.createProductService.Run(body)
	if err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	helpers.ReplyCreated(context, gin.H{"product": product})
}
