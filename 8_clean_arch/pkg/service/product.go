package service

import (
	"go_practice/8_clean_arch/models"
	"go_practice/8_clean_arch/pkg/repository"
)

type ProductSvc struct {
	rep *repository.Repository
}

func NewProductSvc(repo *repository.Repository) *ProductSvc {
	return &ProductSvc{
		rep: repo,
	}
}

func (bsv *ProductSvc) GetProducts() ([]*models.Product, error) {
	return bsv.rep.GetProducts()
}

func (bsv *ProductSvc) GetProductById(id int) (*models.Product, error) {
	return bsv.rep.GetProductById(id)
}

func (bsv *ProductSvc) AddProduct(product models.Product) (int, error) {
	return bsv.rep.AddProduct(product)
}
