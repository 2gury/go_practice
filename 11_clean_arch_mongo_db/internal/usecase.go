package internal

import "go_practice/11_clean_arch_mongo_db/internal/models"

type ProductUsecase interface {
	GetAllProducts() ([]*models.Product, error)
	GetProductById(id string) (*models.Product, error)
	AddProduct(product *models.Product) (string, error)
	UpdateProduct(product *models.Product) (int64, error)
	DeleteProductById(id string) (int64, error)
}
