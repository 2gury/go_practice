package repository

import "go_practice/8_clean_arch/models"

type Repository struct {
	ProductRepository
}

type ProductRepository interface {
	GetProducts() ([]*models.Product, error)
	GetProductById(id int) (*models.Product, error)
	AddProduct(product models.ProductInput) (int, error)
	UpdateProduct(productId int, updatedProduct models.ProductInput) (int, error)
}

func NewRepository(rep *LocalRepository) *Repository {
	return &Repository{
		ProductRepository: NewProductRep(rep),
	}
}
