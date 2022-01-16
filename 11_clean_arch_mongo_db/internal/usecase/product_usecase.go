package usecase

import (
	"go_practice/11_clean_arch_mongo_db/internal"
	"go_practice/11_clean_arch_mongo_db/internal/models"
)

type ProductUsecase struct {
	r internal.ProductRepository
}

func NewProductUsecase(rep internal.ProductRepository) internal.ProductUsecase {
	return &ProductUsecase{
		r: rep,
	}
}

func (u *ProductUsecase) GetAllProducts() ([]*models.Product, error) {
	return u.r.SelectAll()
}

func (u *ProductUsecase) GetProductById(id string) (*models.Product, error) {
	return u.r.SelectById(id)
}

func (u *ProductUsecase) AddProduct(product *models.Product) (string, error) {
	return u.r.Insert(product)
}

func (u *ProductUsecase) UpdateProduct(product *models.Product) (int64, error) {
	return u.r.Update(product)
}

func (u *ProductUsecase) DeleteProductById(id string) (int64, error) {
	return u.r.DeleteById(id)
}
