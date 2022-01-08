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

func (s *ProductSvc) GetProducts() ([]*models.Product, error) {
	return s.rep.GetProducts()
}

func (s *ProductSvc) GetProductById(id int) (*models.Product, error) {
	return s.rep.GetProductById(id)
}

func (s *ProductSvc) AddProduct(product models.ProductInput) (int, error) {
	return s.rep.AddProduct(product)
}

func (s *ProductSvc) UpdateProduct(productId int, updatedProduct models.ProductInput) (int, error) {
	return s.rep.UpdateProduct(productId, updatedProduct)
}
