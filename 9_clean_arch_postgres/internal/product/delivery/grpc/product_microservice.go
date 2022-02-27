package grpc

import (
	"context"
	"database/sql"
	systetmErrors "github.com/pkg/errors"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/product"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct {
	productRep product.ProductRepository
	UnimplementedProductServiceServer
}

func NewProductService(rep product.ProductRepository) *ProductService {
	return &ProductService{
		productRep: rep,
	}
}

func (pm *ProductService) List(ctx context.Context, empty *emptypb.Empty) (*ArrayProducts, error) {
	products, err := pm.productRep.SelectAll()
	if err != nil {
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return ModelProductsToGrpcProducts(products), nil
}

func (pm *ProductService) Create(ctx context.Context, product *Product) (*ProductIdValue, error) {
	if product.Price <= 0 || product.Title == "" {
		return nil, errors.GetErrorFromGrpc(consts.CodeBadRequest, systetmErrors.New(
			"Error when add product. Price should be greater than 0. Title should be not empty"))
	}

	id, err := pm.productRep.Insert(GrpcProductToModel(product))
	if err != nil {
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return &ProductIdValue{Value: id}, nil
}

func (pm *ProductService) GetById(ctx context.Context, id *ProductIdValue) (*Product, error) {
	prod, err := pm.productRep.SelectById(id.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.GetErrorFromGrpc(consts.CodeProductDoesNotExist, err)
		}
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return ModelProductToGrpc(prod), nil
}

func (pm *ProductService) UpdateById(ctx context.Context, updatedProduct *UpdateInfoProduct) (*emptypb.Empty, error) {
	if updatedProduct.Product.Price <= 0 || updatedProduct.Product.Title == "" {
		return &emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeBadRequest, systetmErrors.New(
			"Error when add product. Price should be greater than 0. Title should be not empty"))
	}

	if _, err := pm.GetById(ctx, updatedProduct.Id); err != nil {
		return &emptypb.Empty{}, err
	}
	err := pm.productRep.UpdateById(updatedProduct.Id.Value, GrpcProductToModel(updatedProduct.Product))
	if err != nil {
		return &emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return &emptypb.Empty{}, nil
}

func (pm *ProductService) DeleteById(ctx context.Context, id *ProductIdValue) (*emptypb.Empty, error) {
	if _, err := pm.GetById(ctx, id); err != nil {
		return &emptypb.Empty{}, err
	}
	err := pm.productRep.DeleteById(id.Value)
	if err != nil {
		return &emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return &emptypb.Empty{}, nil
}