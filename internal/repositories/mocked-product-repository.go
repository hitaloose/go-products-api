package repositories

import (
	"time"

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

func (repository MockedProductRepository) Create(product *models.Product) (*models.Product, error) {
	repository.nextId++

	product.ID = repository.nextId
	product.CreatedAt = time.Now().GoString()

	repository.products[product.ID] = product

	return product, nil
}
