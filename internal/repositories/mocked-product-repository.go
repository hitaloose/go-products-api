package repositories

import (
	"time"

	"github.com/hitaloose/go-products-api/internal/dtos"
	"github.com/hitaloose/go-products-api/internal/models"
)

type MockedProductRepository struct {
	products map[int64]*models.Product
	nextId   int64
}

var mockedProductRepository *MockedProductRepository

func NewMockedProductRepository() *MockedProductRepository {
	if mockedProductRepository == nil {
		mockedProductRepository = &MockedProductRepository{
			products: make(map[int64]*models.Product),
			nextId:   0,
		}
	}

	return mockedProductRepository
}

func (repository *MockedProductRepository) Create(product *models.Product) (*models.Product, error) {
	repository.nextId++

	product.ID = repository.nextId
	product.CreatedAt = time.Now().Format(time.RFC3339)

	repository.products[product.ID] = product

	return product, nil
}

func (repository *MockedProductRepository) filterProducts(filters *dtos.ProductFilterDto) []*models.Product {
	var result []*models.Product

	for _, product := range repository.products {
		if filters.Name != nil && *filters.Name != product.Name {
			continue
		}
		if filters.Stock != nil && *filters.Stock != product.Stock {
			continue
		}

		result = append(result, product)
	}

	return result
}

func (repository *MockedProductRepository) GetAll(filters *dtos.ProductFilterDto) ([]*models.Product, error) {
	return repository.filterProducts(filters), nil
}
