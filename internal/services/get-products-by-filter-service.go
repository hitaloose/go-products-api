package services

import (
	"github.com/hitaloose/go-products-api/internal/dtos"
	"github.com/hitaloose/go-products-api/internal/models"
	"github.com/hitaloose/go-products-api/internal/repositories"
)

type GetProductsByFilterService struct {
	productRepository repositories.ProductRepository
}

func NewGetProductsByFilterService() *GetProductsByFilterService {
	return &GetProductsByFilterService{
		productRepository: repositories.NewMockedProductRepository(),
	}
}

func (service GetProductsByFilterService) Run(filters dtos.ProductFilterDto) ([]*models.Product, error) {
	products, err := service.productRepository.GetAll(&filters)
	if err != nil {
		return nil, err
	}

	return products, nil
}
