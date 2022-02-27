package grpc

import (
	"go_practice/9_clean_arch_db/internal/models"
)

func GrpcProductToModel(prod *Product) *models.Product {
	return &models.Product{
		Id: prod.Id,
		Title: prod.Title,
		Price: int(prod.Price),
	}
}

func ModelProductToGrpc(prod *models.Product) *Product {
	return &Product{
		Id: prod.Id,
		Title: prod.Title,
		Price: int64(prod.Price),
	}
}

func ModelProductsToGrpcProducts(products []*models.Product) *ArrayProducts {
	convertedProds :=  &ArrayProducts{}
	for _, product := range products {
		convertedProds.Value = append(convertedProds.Value, ModelProductToGrpc(product))
	}
	return convertedProds
}

func GrpcProductsToModelProducts(products *ArrayProducts) []*models.Product {
	var convertedProds []*models.Product
	for _, product := range products.Value {
		convertedProds = append(convertedProds, GrpcProductToModel(product))
	}
	return convertedProds
}
