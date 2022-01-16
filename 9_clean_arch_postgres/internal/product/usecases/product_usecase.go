package usecases

import (
	"github.com/pkg/errors"
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
)

type ProductUsecase struct {
	productRep product.ProductRepository
}

func NewProductUsecase(rep product.ProductRepository) product.ProductUsecase {
	return &ProductUsecase{
		productRep: rep,
	}
}

func (u *ProductUsecase) List() ([]*models.Product, error) {
	return u.productRep.SelectAll()
}

func (u *ProductUsecase) Create(product models.Product) (uint64, error) {
	if product.Price <= 0 || product.Title == "" {
		return 0, errors.New("Error when add product. Price should be greater than 0")
	}
	return u.productRep.Insert(product)
}
func (u *ProductUsecase) GetById(id uint64) (*models.Product, error) {
	return u.productRep.SelectById(id)
}
func (u *ProductUsecase) UpdateById(productId uint64, updatedProduct models.Product) (bool, error) {
	if updatedProduct.Price <= 0 || updatedProduct.Title == "" {
		return false, errors.New("Error when add product. Price should be greater than 0")
	}
	return u.productRep.UpdateById(productId, updatedProduct)
}

func (u *ProductUsecase) DeleteById(id uint64) (bool, error) {
	return u.productRep.DeleteById(id)
}
