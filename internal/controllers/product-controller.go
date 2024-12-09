package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hitaloose/go-products-api/internal/dtos"
	"github.com/hitaloose/go-products-api/internal/exceptions"
	"github.com/hitaloose/go-products-api/internal/helpers"
	"github.com/hitaloose/go-products-api/internal/services"
)

type ProductController struct {
	createService       services.CreateProductService
	getByFiltersService services.GetProductsByFilterService
}

func NewProductController() *ProductController {
	return &ProductController{
		createService:       *services.NewCreateProductService(),
		getByFiltersService: *services.NewGetProductsByFilterService(),
	}
}

// @Tags Product
// @Router /product  [post]
// @Summary Create a new product
// @ID CreateProduct
// @Param request body dtos.CreateProductDto true "body"
// @Produce json
// @Success 201 {object} dtos.ProductResponseDto
func (controller ProductController) Create(context *gin.Context) {
	var body dtos.CreateProductDto

	if err := helpers.BindJSON(context, &body); err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	if err := helpers.ValidateStruct(&body); err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	product, err := controller.createService.Run(body)
	if err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	helpers.ReplyCreated(context, gin.H{"product": product})
}

// @Tags Product
// @Router /product  [get]
// @Summary Create all products by filters
// @ID GetAllProducts
// @Param title query string false "title"
// @Param stock query string false "stock"
// @Produce json
// @Success 200 {object} dtos.ProductsResponseDto
func (controller ProductController) GetAll(context *gin.Context) {
	var queries dtos.ProductFilterDto

	if err := helpers.BindQuery(context, &queries); err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	if err := helpers.ValidateStruct(&queries); err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	products, err := controller.getByFiltersService.Run(queries)
	if err != nil {
		exceptions.GinExceptionHandler(context, err)
		return
	}

	helpers.ReplyOK(context, gin.H{"products": products})
}
