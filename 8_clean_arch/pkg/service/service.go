package service

import (
	"go_practice/8_clean_arch/models"
	"go_practice/8_clean_arch/pkg/repository"
	"go_practice/8_clean_arch/pkg/service/product"
)

type Service struct {
	ProductService
}

type ProductService interface {
	GetProducts() ([]*models.Product, error)
	GetProductById(id int) (*models.Product, error)
	AddProduct(product models.ProductInput) (int, error)
	UpdateProduct(productId int, updatedProduct models.ProductInput) (int, error)
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		ProductService: product.NewProductSvc(rep),
	}
}
