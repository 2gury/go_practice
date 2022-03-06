package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	mock_product "go_practice/9_clean_arch_db/internal/product/mocks"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestProductService_List(t *testing.T) {
	type mockBehaviour func(productRep *mock_product.MockProductRepository, products *ArrayProducts)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		products      *ArrayProducts
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, products *ArrayProducts) {
				productRep.
					EXPECT().
					SelectAll().
					Return(GrpcProductsToModelProducts(products), nil)
			},
			products: &ArrayProducts{
				Value: []*Product {
					&Product{
						Id:    1,
						Title: "WB",
						Price: 120,
					},
					&Product {
						Id:    2,
						Title: "Ozon",
						Price: 500,
					},
				},
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, products *ArrayProducts) {
				productRep.
					EXPECT().
					SelectAll().
					Return(nil, fmt.Errorf("sql error"))
			},
			products: nil,
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviour(productRep, testCase.products)
			productSvc := NewProductService(productRep)

			prods, err := productSvc.List(context.Background(), &emptypb.Empty{})

			assert.Equal(t, prods, testCase.products)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductService_Create(t *testing.T) {
	type mockBehaviour func(productRep *mock_product.MockProductRepository)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		products      *ArrayProducts
		inProduct     *Product
		expProductId  *ProductIdValue
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(productRep *mock_product.MockProductRepository) {
				productRep.
					EXPECT().
					Insert(gomock.Any()).
					Return(uint64(1), nil)
			},
			inProduct: &Product{
				Title: "WB",
				Price: 120,
			},
			expProductId: &ProductIdValue{Value: 1},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productRep *mock_product.MockProductRepository) {
				productRep.
					EXPECT().
					Insert(gomock.Any()).
					Return(uint64(0), fmt.Errorf("sql error"))
			},
			inProduct: &Product{
				Title: "WB",
				Price: 120,
			},
			expProductId: nil,
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviour(productRep)
			productSvc := NewProductService(productRep)

			lastId, err := productSvc.Create(context.Background(), testCase.inProduct)

			assert.Equal(t, lastId, testCase.expProductId)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductService_GetById(t *testing.T) {
	type mockBehaviour func(productRep *mock_product.MockProductRepository, product *Product)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		products      *ArrayProducts
		inProductId   *ProductIdValue
		outProduct     *Product
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, product *Product) {
				productRep.
					EXPECT().
					SelectById(gomock.Any()).
					Return(GrpcProductToModel(product), nil)
			},
			inProductId: &ProductIdValue{
				Value: 1,
			},
			outProduct: &Product{
				Id: 1,
				Title: "WB",
				Price: 120,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, product *Product) {
				productRep.
					EXPECT().
					SelectById(gomock.Any()).
					Return(nil, fmt.Errorf("sql error"))
			},
			inProductId: &ProductIdValue{
				Value: 1,
			},
			outProduct: nil,
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
		{
			name: "Error: CodeProductDoesNotExist",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, product *Product) {
				productRep.
					EXPECT().
					SelectById(gomock.Any()).
					Return(nil, sql.ErrNoRows)
			},
			inProductId: &ProductIdValue{
				Value: 1,
			},
			outProduct: nil,
			expError: errors.GetErrorFromGrpc(consts.CodeProductDoesNotExist, fmt.Errorf("sql: no rows in result set")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviour(productRep, testCase.outProduct)
			productSvc := NewProductService(productRep)

			product, err := productSvc.GetById(context.Background(), testCase.inProductId)

			assert.Equal(t, product, testCase.outProduct)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductService_UpdateById(t *testing.T) {
	type mockBehaviourSelectById func(productRep *mock_product.MockProductRepository, updInfo *UpdateInfoProduct)
	type mockBehaviourUpdateById func(productRep *mock_product.MockProductRepository, updInfo *UpdateInfoProduct)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviourSelectById mockBehaviourSelectById
		mockBehaviourUpdateById mockBehaviourUpdateById
		products      *ArrayProducts
		inUpdateInfoProduct   *UpdateInfoProduct
		expError      error
	}{
		{
			name: "OK",
			mockBehaviourSelectById: func(productRep *mock_product.MockProductRepository, updInfo *UpdateInfoProduct) {
				productRep.
					EXPECT().
					SelectById(updInfo.Id.Value).
					Return(GrpcProductToModel(updInfo.Product), nil)
			},
			mockBehaviourUpdateById: func(productRep *mock_product.MockProductRepository, updInfo *UpdateInfoProduct) {
				productRep.
					EXPECT().
					UpdateById(updInfo.Id.Value, GrpcProductToModel(updInfo.Product)).
					Return(nil)
			},
			inUpdateInfoProduct: &UpdateInfoProduct{
				Id: &ProductIdValue{
					Value: 1,
				},
				Product: &Product{
					Title: "OZON",
					Price: 150,
				},
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviourSelectById: func(productRep *mock_product.MockProductRepository, updInfo *UpdateInfoProduct) {
				productRep.
					EXPECT().
					SelectById(updInfo.Id.Value).
					Return(GrpcProductToModel(updInfo.Product), nil)
			},
			mockBehaviourUpdateById: func(productRep *mock_product.MockProductRepository, updInfo *UpdateInfoProduct) {
				productRep.
					EXPECT().
					UpdateById(updInfo.Id.Value, GrpcProductToModel(updInfo.Product)).
					Return(fmt.Errorf("sql error"))
			},
			inUpdateInfoProduct: &UpdateInfoProduct{
				Id: &ProductIdValue{
					Value: 1,
				},
				Product: &Product{
					Title: "OZON",
					Price: 150,
				},
			},
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
		{
			name: "Error: CodeProductDoesNotExist",
			mockBehaviourSelectById: func(productRep *mock_product.MockProductRepository, updInfo *UpdateInfoProduct) {
				productRep.
					EXPECT().
					SelectById(updInfo.Id.Value).
					Return(nil, sql.ErrNoRows)
			},
			mockBehaviourUpdateById: func(productRep *mock_product.MockProductRepository, updInfo *UpdateInfoProduct) {},
			inUpdateInfoProduct: &UpdateInfoProduct{
				Id: &ProductIdValue{
					Value: 1,
				},
				Product: &Product{
					Title: "OZON",
					Price: 150,
				},
			},
			expError: errors.GetErrorFromGrpc(consts.CodeProductDoesNotExist, fmt.Errorf("sql: no rows in result set")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviourSelectById(productRep, testCase.inUpdateInfoProduct)
			testCase.mockBehaviourUpdateById(productRep, testCase.inUpdateInfoProduct)
			productSvc := NewProductService(productRep)

			_, err := productSvc.UpdateById(context.Background(), testCase.inUpdateInfoProduct)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductService_DeleteById(t *testing.T) {
	type mockBehaviourSelectById func(productRep *mock_product.MockProductRepository, productId *ProductIdValue)
	type mockBehaviourDeleteById func(productRep *mock_product.MockProductRepository, productId *ProductIdValue)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviourSelectById mockBehaviourSelectById
		mockBehaviourDeleteById mockBehaviourDeleteById
		products      *ArrayProducts
		inProductId   *ProductIdValue
		expError      error
	}{
		{
			name: "OK",
			mockBehaviourSelectById: func(productRep *mock_product.MockProductRepository, productId *ProductIdValue) {
				productRep.
					EXPECT().
					SelectById(productId.Value).
					Return(&models.Product{}, nil)
			},
			mockBehaviourDeleteById: func(productRep *mock_product.MockProductRepository, productId *ProductIdValue) {
				productRep.
					EXPECT().
					DeleteById(productId.Value).
					Return(nil)
			},
			inProductId: &ProductIdValue{
				Value: 1,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviourSelectById: func(productRep *mock_product.MockProductRepository, productId *ProductIdValue) {
				productRep.
					EXPECT().
					SelectById(productId.Value).
					Return(&models.Product{}, nil)
			},
			mockBehaviourDeleteById: func(productRep *mock_product.MockProductRepository, productId *ProductIdValue) {
				productRep.
					EXPECT().
					DeleteById(productId.Value).
					Return(fmt.Errorf("sql error"))
			},
			inProductId: &ProductIdValue{
				Value: 1,
			},
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviourSelectById(productRep, testCase.inProductId)
			testCase.mockBehaviourDeleteById(productRep, testCase.inProductId)
			productSvc := NewProductService(productRep)

			_, err := productSvc.DeleteById(context.Background(), testCase.inProductId)

			assert.Equal(t, err, testCase.expError)
		})
	}
}
