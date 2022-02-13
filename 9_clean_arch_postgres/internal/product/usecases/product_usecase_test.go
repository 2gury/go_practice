package usecases

import (
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
					Return(nil, fmt.Errorf("SQL Error"))
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