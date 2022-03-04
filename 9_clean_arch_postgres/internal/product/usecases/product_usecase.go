package usecases

import (
	"context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/product"
	"go_practice/9_clean_arch_db/internal/product/delivery/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductUsecase struct {
	productSvc grpc.ProductServiceClient
}

func NewProductUsecase(svc grpc.ProductServiceClient) product.ProductUsecase {
	return &ProductUsecase{
		productSvc: svc,
	}
}

func (u *ProductUsecase) List() ([]*models.Product, *errors.Error) {
	products, err := u.productSvc.List(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, errors.GetCustomError(err)
	}

	return grpc.GrpcProductsToModelProducts(products), nil
}

func (u *ProductUsecase) Create(product *models.Product) (uint64, *errors.Error) {
	id, err := u.productSvc.Create(context.Background(), grpc.ModelProductToGrpc(product))
	if err != nil {
		return 0, errors.GetCustomError(err)
	}

	return id.Value, nil
}
func (u *ProductUsecase) GetById(id uint64) (*models.Product, *errors.Error) {
	prod, err := u.productSvc.GetById(context.Background(), &grpc.ProductIdValue{Value: id})
	if err != nil {
		return nil, errors.GetCustomError(err)
	}

	return grpc.GrpcProductToModel(prod), nil
}
func (u *ProductUsecase) UpdateById(productId uint64, updatedProduct *models.Product) *errors.Error {
	_, err := u.productSvc.UpdateById(context.Background(), &grpc.UpdateInfoProduct{
		Id: &grpc.ProductIdValue{Value: productId},
		Product: grpc.ModelProductToGrpc(updatedProduct),
	})
	if err != nil {
		return errors.GetCustomError(err)
	}

	return nil
}

func (u *ProductUsecase) DeleteById(id uint64) *errors.Error {
	_, err := u.productSvc.DeleteById(context.Background(), &grpc.ProductIdValue{Value: id})
	if err != nil {
		return errors.GetCustomError(err)
	}

	return nil
}
