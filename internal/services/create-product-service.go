package services

import (
	"github.com/hitaloose/go-products-api/internal/dtos"
	"github.com/hitaloose/go-products-api/internal/models"
	"github.com/hitaloose/go-products-api/internal/repositories"
)

type CreateProductService struct {
	productRepository repositories.ProductRepository
}

func NewCreateProductService() *CreateProductService {
	return &CreateProductService{
		productRepository: repositories.NewMockedProductRepository(),
	}
}

func (service CreateProductService) Run(dto dtos.CreateProductDto) (*models.Product, error) {
	product, err := service.productRepository.Create(&models.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
		Price:       dto.Price,
		Stock:       dto.Stock,
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}
