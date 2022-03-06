package usecases

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/product/delivery/grpc"
	mock_grpc "go_practice/9_clean_arch_db/internal/product/delivery/grpc/mocks"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestProductUsecase_List(t *testing.T) {
	type mockBehaviour func(productSvc *mock_grpc.MockProductServiceClient, products []*models.Product)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		products      []*models.Product
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(productSvc *mock_grpc.MockProductServiceClient, products []*models.Product) {
				productSvc.
					EXPECT().
					List(context.Background(), &emptypb.Empty{}).
					Return(grpc.ModelProductsToGrpcProducts(products), nil)
			},
			products: []*models.Product {
				&models.Product{
					Id:    1,
					Title: "WB",
					Price: 120,
				},
				&models.Product {
					Id:    2,
					Title: "Ozon",
					Price: 500,
				},
			},
			expError: (*errors.Error)(nil),
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productSvc *mock_grpc.MockProductServiceClient, products []*models.Product) {
				productSvc.
					EXPECT().
					List(context.Background(), &emptypb.Empty{}).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			products: nil,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productSvc := mock_grpc.NewMockProductServiceClient(ctrl)
			testCase.mockBehaviour(productSvc, testCase.products)
			productUsecase := NewProductUsecase(productSvc)

			prods, err := productUsecase.List()

			assert.Equal(t, prods, testCase.products)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductUsecase_Create(t *testing.T) {
	type mockBehaviour func(productSvc *mock_grpc.MockProductServiceClient, product *models.Product)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		product       *models.Product
		expError      *errors.Error
		expOutput     uint64
	}{
		{
			name: "OK",
			mockBehaviour: func(productSvc *mock_grpc.MockProductServiceClient, product *models.Product) {
				productSvc.
					EXPECT().
					Create(context.Background(), grpc.ModelProductToGrpc(product)).
					Return(&grpc.ProductIdValue{Value: 1}, nil)
			},
			product: &models.Product{
				Title: "WB",
				Price: 120,
			},
			expError:  nil,
			expOutput: 1,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productSvc *mock_grpc.MockProductServiceClient, product *models.Product) {
				productSvc.
					EXPECT().
					Create(context.Background(), grpc.ModelProductToGrpc(product)).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			product: &models.Product{
				Title: "WB",
				Price: 120,
			},
			expError:  errors.Get(consts.CodeInternalError),
			expOutput: 0,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productSvc := mock_grpc.NewMockProductServiceClient(ctrl)
			testCase.mockBehaviour(productSvc, testCase.product)
			productUsecase := NewProductUsecase(productSvc)

			lastId, err := productUsecase.Create(testCase.product)

			assert.Equal(t, lastId, testCase.expOutput)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductUsecase_GetById(t *testing.T) {
	type mockBehaviour func(productSvc *mock_grpc.MockProductServiceClient, productID uint64, product *models.Product)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		productId     uint64
		expError      *errors.Error
		expOutput     *models.Product
	}{
		{
			name: "OK",
			mockBehaviour: func(productSvc *mock_grpc.MockProductServiceClient, productId uint64, product *models.Product) {
				productSvc.
					EXPECT().
					GetById(context.Background(), &grpc.ProductIdValue{Value: productId}).
					Return(grpc.ModelProductToGrpc(product), nil)
			},
			productId: 1,
			expError:  nil,
			expOutput: &models.Product{
				Id:    1,
				Title: "WB",
				Price: 120,
			},
		},
		{
			name: "Error: CodeProductDoesNotExist",
			mockBehaviour: func(productSvc *mock_grpc.MockProductServiceClient, productId uint64, product *models.Product) {
				productSvc.
					EXPECT().
					GetById(context.Background(), &grpc.ProductIdValue{Value: productId}).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeProductDoesNotExist, errors.NilErrror))
			},
			productId: 1000,
			expError:  errors.Get(consts.CodeProductDoesNotExist),
			expOutput: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productSvc *mock_grpc.MockProductServiceClient, productId uint64, product *models.Product) {
				productSvc.
					EXPECT().
					GetById(context.Background(), &grpc.ProductIdValue{Value: productId}).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			productId: 1,
			expError:  errors.Get(consts.CodeInternalError),
			expOutput: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productSvc := mock_grpc.NewMockProductServiceClient(ctrl)
			testCase.mockBehaviour(productSvc, testCase.productId, testCase.expOutput)
			productUsecase := NewProductUsecase(productSvc)

			lastId, err := productUsecase.GetById(testCase.productId)

			assert.Equal(t, lastId, testCase.expOutput)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductUsecase_UpdateById(t *testing.T) {
	type mockBehaviourGetById func(productSvc *mock_grpc.MockProductServiceClient, productId uint64)
	type mockBehaviourUpdateById func(productSvc *mock_grpc.MockProductServiceClient, productID uint64, product *models.Product)
	t.Parallel()

	testTable := []struct {
		name                    string
		mockBehaviourUpdateById mockBehaviourUpdateById
		productId               uint64
		product                 *models.Product
		expError                *errors.Error
	}{
		{
			name: "OK",
			mockBehaviourUpdateById: func(productSvc *mock_grpc.MockProductServiceClient, productId uint64, product *models.Product) {
				productSvc.
					EXPECT().
					UpdateById(context.Background(), &grpc.UpdateInfoProduct{
						Id: &grpc.ProductIdValue{Value: productId},
						Product: grpc.ModelProductToGrpc(product),
					}).
					Return(&emptypb.Empty{}, nil)
			},
			productId: 1,
			product: &models.Product{
				Title: "Ozon",
				Price: 500,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviourUpdateById: func(productSvc *mock_grpc.MockProductServiceClient, productId uint64, product *models.Product) {
				productSvc.
					EXPECT().
					UpdateById(context.Background(), &grpc.UpdateInfoProduct{
						Id: &grpc.ProductIdValue{Value: productId},
						Product: grpc.ModelProductToGrpc(product),
					}).
					Return(&emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			productId: 1,
			product: &models.Product{
				Title: "Ozon",
				Price: 500,
			},
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productSvc := mock_grpc.NewMockProductServiceClient(ctrl)
			testCase.mockBehaviourUpdateById(productSvc, testCase.productId, testCase.product)
			productUsecase := NewProductUsecase(productSvc)

			err := productUsecase.UpdateById(testCase.productId, testCase.product)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductUsecase_DeleteById(t *testing.T) {
	type mockBehaviourGetById func(productSvc *mock_grpc.MockProductServiceClient, productId uint64)
	type mockBehaviourDeleteById func(productSvc *mock_grpc.MockProductServiceClient, productID uint64)
	t.Parallel()

	testTable := []struct {
		name                    string
		mockBehaviourDeleteById mockBehaviourDeleteById
		productId               uint64
		product                 *models.Product
		expError                *errors.Error
	}{
		{
			name: "OK",
			mockBehaviourDeleteById: func(productSvc *mock_grpc.MockProductServiceClient, productId uint64) {
				productSvc.
					EXPECT().
					DeleteById(context.Background(), &grpc.ProductIdValue{Value: productId}).
					Return(&emptypb.Empty{}, nil)
			},
			productId: 1,
			expError:  nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviourDeleteById: func(productSvc *mock_grpc.MockProductServiceClient, productId uint64) {
				productSvc.
					EXPECT().
					DeleteById(context.Background(), &grpc.ProductIdValue{Value: productId}).
					Return(&emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			productId: 1,
			expError:  errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productSvc := mock_grpc.NewMockProductServiceClient(ctrl)
			testCase.mockBehaviourDeleteById(productSvc, testCase.productId)
			productUsecase := NewProductUsecase(productSvc)

			err := productUsecase.DeleteById(testCase.productId)

			assert.Equal(t, err, testCase.expError)
		})
	}
}
