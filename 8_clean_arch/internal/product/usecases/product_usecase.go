package usecases

import (
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
)

type ProductUsecase struct {
	r product.ProductRepository
}

func NewProductUsecase(rep product.ProductRepository) product.ProductUsecase {
	return &ProductUsecase{
		r: rep,
	}
}

func (u *ProductUsecase) List() ([]*models.Product, error) {
	return u.r.SelectAll()
}

func (u *ProductUsecase) Create(product models.ProductInput) (uint64, error) {
	return u.r.Insert(product)
}
func (u *ProductUsecase) GetById(id uint64) (*models.Product, error) {
	return u.r.SelectById(id)
}
func (u *ProductUsecase) UpdateById(productId uint64, updatedProduct models.ProductInput) (int, error) {
	return u.r.UpdateById(productId, updatedProduct)
}