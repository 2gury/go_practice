package repository

import (
	"github.com/pkg/errors"
	"go_practice/8_clean_arch/models"
)

type ProductRep struct {
	data *LocalRepository
}

func NewProductRep(mapRep *LocalRepository) *ProductRep {
	return &ProductRep{
		data: mapRep,
	}
}

func (r *ProductRep) GetProducts() ([]*models.Product, error) {
	return r.data.Products, nil
}

func (r *ProductRep) GetProductById(id int) (*models.Product, error) {
	for _, product := range r.data.Products {
		if product.Id == id {
			return product, nil
		}
	}
	return nil, nil
}

func (r *ProductRep) AddProduct(product models.ProductInput) (int, error) {
	if product.Price <= 0 || product.Title == "" {
		return -1, errors.New("Error when add product. Price should be greater than 0")
	}
	r.data.mu.Lock()
	defer r.data.mu.Unlock()
	maxId := 0
	for i, product := range r.data.Products {
		if product.Id > maxId {
			maxId = i
		}
	}
	maxId++
	r.data.Products = append(r.data.Products, &models.Product{
		Id:    maxId,
		Title: product.Title,
		Price: product.Price,
	})
	return maxId, nil
}

func (r *ProductRep) UpdateProduct(productId int, updatedProduct models.ProductInput) (int, error) {
	if updatedProduct.Price <= 0 || updatedProduct.Title == "" {
		return 0, errors.New("Error when add product. Price should be greater than 0")
	}
	product, err := r.GetProductById(productId)
	if product == nil {
		return 0, errors.New("Error when add product. No product with this id")
	}
	if err != nil {
		return 0, errors.New("Error when add product")
	}
	r.data.mu.Lock()
	defer r.data.mu.Unlock()
	product.Title = updatedProduct.Title
	product.Price = updatedProduct.Price
	return 1, nil
}
