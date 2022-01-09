package repository

import (
	"github.com/pkg/errors"
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
	"sync"
)

type ProductArrayRepository struct {
	Products []*models.Product
	mu       *sync.Mutex
}

func NewProductArrayRepository() product.ProductRepository {
	return &ProductArrayRepository{
		Products: []*models.Product{},
		mu: &sync.Mutex{},
	}
}

func (r *ProductArrayRepository) SelectAll() ([]*models.Product, error) {
	return r.Products, nil
}

func (r *ProductArrayRepository) SelectById(id uint64) (*models.Product, error) {
	for _, product := range r.Products {
		if product.Id == id {
			return product, nil
		}
	}
	return nil, nil
}

func (r *ProductArrayRepository) Insert(product models.ProductInput) (uint64, error) {
	if product.Price <= 0 || product.Title == "" {
		return 0, errors.New("Error when add product. Price should be greater than 0")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	var maxId uint64 = 0
	for i, product := range r.Products {
		if product.Id > maxId {
			maxId = uint64(i)
		}
	}
	maxId++
	r.Products = append(r.Products, &models.Product{
		Id:    maxId,
		Title: product.Title,
		Price: product.Price,
	})
	return maxId, nil
}

func (r *ProductArrayRepository) UpdateById(productId uint64, updatedProduct models.ProductInput) (int, error) {
	if updatedProduct.Price <= 0 || updatedProduct.Title == "" {
		return 0, errors.New("Error when add product. Price should be greater than 0")
	}
	product, err := r.SelectById(productId)
	if product == nil {
		return 0, errors.New("Error when add product. No product with this id")
	}
	if err != nil {
		return 0, errors.New("Error when add product")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	product.Title = updatedProduct.Title
	product.Price = updatedProduct.Price
	return 1, nil
}
