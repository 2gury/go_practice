package internal

import "go_practice/11_clean_arch_mongo_db/internal/models"

type ProductRepository interface {
	SelectAll() ([]*models.Product, error)
	SelectById(id string) (*models.Product, error)
	Insert(product *models.Product) (string, error)
	Update(product *models.Product) (int64, error)
	DeleteById(id string) (int64, error)
}
