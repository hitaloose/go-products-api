package dtos

import "github.com/hitaloose/go-products-api/internal/models"

type ProductResponseDto struct {
	Product *models.Product `json:"product"`
}

type ProductsResponseDto struct {
	Products []*models.Product `json:"products"`
}
