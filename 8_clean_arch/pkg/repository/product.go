package repository

import "go_practice/8_clean_arch/models"

type ProductRep struct {
	rep *InRuntimeRepository
}

func NewProductRep(mapRep *InRuntimeRepository) *ProductRep {
	return &ProductRep{
		rep: mapRep,
	}
}

func (brp *ProductRep) GetProducts() ([]*models.Product, error) {
	return brp.rep.Products, nil
}

func (brp *ProductRep) GetProductById(id int) (*models.Product, error) {
	for i := 0; i < len(brp.rep.Products); i++ {
		if brp.rep.Products[i].Id == id {
			return brp.rep.Products[i], nil
		}
	}
	return nil, nil
}

func (brp *ProductRep) AddProduct(product models.Product) (int, error) {
	brp.rep.mu.Lock()
	defer brp.rep.mu.Unlock()
	brp.rep.Products = append(brp.rep.Products, &models.Product{
		Id: brp.rep.nextId,
		Title: product.Title,
		Price: product.Price,
	})
	brp.rep.nextId++
	return brp.rep.nextId+ - 1, nil
}

