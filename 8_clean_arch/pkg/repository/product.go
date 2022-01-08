package repository

import "go_practice/8_clean_arch/models"

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
	for i := 0; i < len(r.data.Products); i++ {
		if r.data.Products[i].Id == id {
			return r.data.Products[i], nil
		}
	}
	return nil, nil
}

func (r *ProductRep) AddProduct(product models.Product) (int, error) {
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

