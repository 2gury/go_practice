package usecases

import (
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	mock_product "go_practice/9_clean_arch_db/internal/product/mocks"
	"testing"
)

func TestProductHandler_GetProducts(t *testing.T) {
	type mockBehaviour func(productRep *mock_product.MockProductRepository, products []*models.Product)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		products []*models.Product
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, products []*models.Product) {
				productRep.
					EXPECT().
					SelectAll().
					Return(products, nil)
			},
			products: []*models.Product {
				&models.Product{
					Id: 1,
					Title: "WB",
					Price: 120,
				},
				&models.Product{
					Id: 2,
					Title: "Ozon",
					Price: 500,
				},
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, products []*models.Product)  {
				productRep.
					EXPECT().
					SelectAll().
					Return(nil, fmt.Errorf("repository error"))
			},
			products: nil,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviour(productRep, testCase.products)
			productUsecase := NewProductUsecase(productRep)

			prods, err := productUsecase.List()

			assert.Equal(t, prods, testCase.products)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductUsecase_Create(t *testing.T) {
	type mockBehaviour func(productRep *mock_product.MockProductRepository, product *models.Product)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		product *models.Product
		expError *errors.Error
		expOutput uint64
	}{
		{
			name: "OK",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, product *models.Product) {
				productRep.
					EXPECT().
					Insert(product).
					Return(uint64(1), nil)
			},
			product: &models.Product{
				Title: "WB",
				Price: 120,
			},
			expError: nil,
			expOutput: 1,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, product *models.Product) {
				productRep.
					EXPECT().
					Insert(product).
					Return(uint64(0), fmt.Errorf("repository error"))
			},
			product: &models.Product{
				Title: "WB",
				Price: 120,
			},
			expError: errors.Get(consts.CodeInternalError),
			expOutput: 0,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviour(productRep, testCase.product)
			productUsecase := NewProductUsecase(productRep)

			lastId, err := productUsecase.Create(testCase.product)

			assert.Equal(t, lastId, testCase.expOutput)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductUsecase_GetById(t *testing.T) {
	type mockBehaviour func(productRep *mock_product.MockProductRepository, productID uint64, product *models.Product)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		productId uint64
		expError *errors.Error
		expOutput *models.Product
	}{
		{
			name: "OK",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, productId uint64, product *models.Product) {
				productRep.
					EXPECT().
					SelectById(productId).
					Return(product, nil)
			},
			productId: 1,
			expError: nil,
			expOutput: &models.Product{
				Id: 1,
				Title: "WB",
				Price: 120,
			},
		},
		{
			name: "Error: CodeProductDoesNotExist",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, productId uint64, product *models.Product) {
				productRep.
					EXPECT().
					SelectById(productId).
					Return(nil, sql.ErrNoRows)
			},
			productId: 1000,
			expError: errors.Get(consts.CodeProductDoesNotExist),
			expOutput: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productRep *mock_product.MockProductRepository, productId uint64, product *models.Product) {
				productRep.
					EXPECT().
					SelectById(productId).
					Return(nil, fmt.Errorf("repository error"))
			},
			productId: 1,
			expError: errors.Get(consts.CodeInternalError),
			expOutput: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviour(productRep, testCase.productId, testCase.expOutput)
			productUsecase := NewProductUsecase(productRep)

			lastId, err := productUsecase.GetById(testCase.productId)

			assert.Equal(t, lastId, testCase.expOutput)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductUsecase_UpdateById(t *testing.T) {
	type mockBehaviourGetById func(productRep *mock_product.MockProductRepository, productId uint64)
	type mockBehaviourUpdateById func(productRep *mock_product.MockProductRepository, productID uint64, product *models.Product)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviourGetById mockBehaviourGetById
		mockBehaviourUpdateById mockBehaviourUpdateById
		productId uint64
		product *models.Product
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviourGetById: func(productRep *mock_product.MockProductRepository, productId uint64) {
				productRep.
					EXPECT().
					SelectById(productId).
					Return(nil ,nil)
			},
			mockBehaviourUpdateById: func(productRep *mock_product.MockProductRepository, productId uint64, product *models.Product) {
				productRep.
					EXPECT().
					UpdateById(productId, product).
					Return(nil)
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
			mockBehaviourGetById: func(productRep *mock_product.MockProductRepository, productId uint64) {
				productRep.
					EXPECT().
					SelectById(productId).
					Return(nil ,nil)
			},
			mockBehaviourUpdateById: func(productRep *mock_product.MockProductRepository, productId uint64, product *models.Product) {
				productRep.
					EXPECT().
					UpdateById(productId, product).
					Return(fmt.Errorf("repository error"))
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
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviourGetById(productRep, testCase.productId)
			testCase.mockBehaviourUpdateById(productRep, testCase.productId, testCase.product)
			productUsecase := NewProductUsecase(productRep)

			err := productUsecase.UpdateById(testCase.productId, testCase.product)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestProductUsecase_DeleteById(t *testing.T) {
	type mockBehaviourGetById func(productRep *mock_product.MockProductRepository, productId uint64)
	type mockBehaviourDeleteById func(productRep *mock_product.MockProductRepository, productID uint64)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviourGetById mockBehaviourGetById
		mockBehaviourDeleteById mockBehaviourDeleteById
		productId uint64
		product *models.Product
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviourGetById: func(productRep *mock_product.MockProductRepository, productId uint64) {
				productRep.
					EXPECT().
					SelectById(productId).
					Return(nil ,nil)
			},
			mockBehaviourDeleteById: func(productRep *mock_product.MockProductRepository, productId uint64) {
				productRep.
					EXPECT().
					DeleteById(productId).
					Return(nil)
			},
			productId: 1,
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviourGetById: func(productRep *mock_product.MockProductRepository, productId uint64) {
				productRep.
					EXPECT().
					SelectById(productId).
					Return(nil ,nil)
			},
			mockBehaviourDeleteById: func(productRep *mock_product.MockProductRepository, productId uint64) {
				productRep.
					EXPECT().
					DeleteById(productId).
					Return(fmt.Errorf("repository error"))
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
			productRep := mock_product.NewMockProductRepository(ctrl)
			testCase.mockBehaviourGetById(productRep, testCase.productId)
			testCase.mockBehaviourDeleteById(productRep, testCase.productId)
			productUsecase := NewProductUsecase(productRep)

			err := productUsecase.DeleteById(testCase.productId)

			assert.Equal(t, err, testCase.expError)
		})
	}
}