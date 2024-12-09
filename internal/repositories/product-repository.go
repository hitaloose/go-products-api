package repositories

import (
	"github.com/hitaloose/go-products-api/internal/dtos"
	"github.com/hitaloose/go-products-api/internal/models"
)

type ProductRepository interface {
	Create(product *models.Product) (*models.Product, error)
	GetAll(filters *dtos.ProductFilterDto) ([]*models.Product, error)
}
