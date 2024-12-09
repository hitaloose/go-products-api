package services

import (
	"testing"

	"github.com/hitaloose/go-products-api/internal/dtos"
)

func TestIfProductIsCreatedWithCorrectValues(t *testing.T) {
	sut := NewCreateProductService()

	name := "product title"
	description := "product description"
	price := 10.0
	stock := 10
	category := "product category"

	product, _ := sut.Run(dtos.CreateProductDto{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Category:    category,
	})

	if product.Name != name {
		t.Errorf("product.Name")
	}
	if product.Description != description {
		t.Errorf("product.Description")
	}
	if product.Price != price {
		t.Errorf("product.Price")
	}
	if product.Stock != stock {
		t.Errorf("product.Stock")
	}
	if product.Category != category {
		t.Errorf("product.Category")
	}
}
