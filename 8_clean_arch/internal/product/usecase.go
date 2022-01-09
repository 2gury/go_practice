package product

import (
	"go_practice/8_clean_arch/internal/models"
)

type ProductUsecase interface {
	List() ([]*models.Product, error)
	Create(product models.ProductInput) (uint64, error)
	GetById(id uint64) (*models.Product, error)
	UpdateById(productId uint64, updatedProduct models.ProductInput) (int, error)
}
