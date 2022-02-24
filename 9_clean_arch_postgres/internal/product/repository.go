package product

import "go_practice/9_clean_arch_db/internal/models"

type ProductRepository interface {
	SelectAll() ([]*models.Product, error)
	SelectById(id uint64) (*models.Product, error)
	Insert(product *models.Product) (uint64, error)
	UpdateById(productId uint64, updatedProduct *models.Product) error
	DeleteById(id uint64) error
}
