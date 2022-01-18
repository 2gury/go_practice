package product

import (
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/helpers/errors"
)

type ProductUsecase interface {
	List() ([]*models.Product, *errors.Error)
	Create(product models.Product) (uint64, *errors.Error)
	GetById(id uint64) (*models.Product, *errors.Error)
	UpdateById(productId uint64, updatedProduct models.Product) *errors.Error
	DeleteById(id uint64) *errors.Error
}
