package delivery

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	mock_product "go_practice/9_clean_arch_db/internal/product/mocks"
	"go_practice/9_clean_arch_db/tools/converter"
	"go_practice/9_clean_arch_db/tools/response"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestProductHandler_GetProducts(t *testing.T) {
	type mockBehaviour func(s *mock_product.MockProductUsecase)
	t.Parallel()

	products := []*models.Product {
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
	}

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inPath string
		expStatusCode int
		expRespBody response.Response
	}{
		{
			name: "OK",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase) {
				productUse.
					EXPECT().
					List().
					Return(products, nil)
			},
			inPath: "/api/v1/product",
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"products": products,
				},
			},
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase) {
				productUse.
					EXPECT().
					List().
					Return(nil, errors.Get(consts.CodeInternalError))
			},
			inPath: "/api/v1/product",
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("GET", testCase.inPath, nil)
			w := httptest.NewRecorder()
			productUse := mock_product.NewMockProductUsecase(ctrl)
			testCase.mockBehaviour(productUse)
			productHandler := NewProductHandler(productUse)
			productHandler.Configure(mx, nil)

			productHandler.GetProducts()(w, r)
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}

func TestProductHandler_GetProductById(t *testing.T) {
	type mockBehaviour func(s *mock_product.MockProductUsecase, prodId uint64)
	t.Parallel()

	product := &models.Product {
		Id: 1,
		Title: "WB",
		Price: 120,
	}

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inPath string
		inParams map[string]string
		expStatusCode int
		expRespBody response.Response
	}{
		{
			name: "OK",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64) {
				productUse.
					EXPECT().
					GetById(prodId).
					Return(product, nil)
			},
			inPath: "/api/v1/product/1",
			inParams: map[string]string {
				"id": "1",
			},
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"product": product,
				},
			},
		},
		{
			name: "Error: CodeProductDoesNotExist",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64) {
				productUse.
					EXPECT().
					GetById(prodId).
					Return(nil, errors.Get(consts.CodeProductDoesNotExist))
			},
			inPath: "/api/v1/product/1000",
			inParams: map[string]string {
				"id": "1000",
			},
			expStatusCode: errors.Get(consts.CodeProductDoesNotExist).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeProductDoesNotExist),
			},
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64) {
				productUse.
					EXPECT().
					GetById(prodId).
					Return(nil, errors.Get(consts.CodeInternalError))
			},
			inPath: "/api/v1/product/1",
			inParams: map[string]string {
				"id": "1",
			},
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("GET", testCase.inPath, nil)
			r  = mux.SetURLVars(r, testCase.inParams)
			w := httptest.NewRecorder()
			productUse := mock_product.NewMockProductUsecase(ctrl)

			productId, _ := mux.Vars(r)["id"]
			intProductId, _ := strconv.ParseUint(productId, 10, 64)
			testCase.mockBehaviour(productUse, intProductId)
			productHandler := NewProductHandler(productUse)
			productHandler.Configure(mx, nil)

			productHandler.GetProductById()(w, r)
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}

func TestProductHandler_AddProduct(t *testing.T) {
	type mockBehaviour func(productUse *mock_product.MockProductUsecase, prod *models.Product)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inPath string
		inProduct *models.Product
		expStatusCode int
		expRespBody response.Response
	}{
		{
			name: "OK",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prod *models.Product) {
				productUse.
					EXPECT().
					Create(prod).
					Return(uint64(1), nil)
			},
			inPath: "/api/v1/product/",
			inProduct: &models.Product {
				Title: "WB",
				Price: 120,
			},
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"id": 1,
				},
			},
		},
		{
			name: "Error: CodeBadRequest",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prod *models.Product) {
				productUse.
					EXPECT().
					Create(prod).
					Return(uint64(0), errors.Get(consts.CodeBadRequest))
			},
			inPath: "/api/v1/product/",
			inProduct: &models.Product {
				Title: "WB",
				Price: -100,
			},
			expStatusCode: errors.Get(consts.CodeBadRequest).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeBadRequest),
			},
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prod *models.Product) {
				productUse.
					EXPECT().
					Create(prod).
					Return(uint64(0), errors.Get(consts.CodeInternalError))
			},
			inPath: "/api/v1/product/",
			inProduct: &models.Product {
				Title: "Ozon",
				Price: 500,
			},
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("PUT", testCase.inPath, converter.AnyBytesToString(testCase.inProduct))
			w := httptest.NewRecorder()
			productUse := mock_product.NewMockProductUsecase(ctrl)

			testCase.mockBehaviour(productUse, testCase.inProduct)
			productHandler := NewProductHandler(productUse)
			productHandler.Configure(mx, nil)

			productHandler.AddProduct()(w, r)
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}

func TestProductHandler_UpdateProductById(t *testing.T) {
	type mockBehaviour func(productUse *mock_product.MockProductUsecase, prodId uint64, prod *models.Product)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inPath string
		inParams map[string]string
		inProduct *models.Product
		expStatusCode int
		expRespBody response.Response
	}{
		{
			name: "OK",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64, prod *models.Product) {
				productUse.
					EXPECT().
					UpdateById(prodId, prod).
					Return(nil)
			},
			inPath: "/api/v1/product/1",
			inParams: map[string]string {
				"id": "1",
			},
			inProduct: &models.Product {
				Title: "WB",
				Price: 120,
			},
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"updated_elements": true,
				},
			},
		},
		{
			name: "Error: CodeProductDoesNotExist",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64, prod *models.Product) {
				productUse.
					EXPECT().
					UpdateById(prodId, prod).
					Return(errors.Get(consts.CodeProductDoesNotExist))
			},
			inPath: "/api/v1/product/1000",
			inParams: map[string]string {
				"id": "1000",
			},
			inProduct: &models.Product {
				Title: "WB",
				Price: 500,
			},
			expStatusCode: errors.Get(consts.CodeProductDoesNotExist).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeProductDoesNotExist),
			},
		},
		{
			name: "Error: CodeBadRequest",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64, prod *models.Product) {
				productUse.
					EXPECT().
					UpdateById(prodId, prod).
					Return(errors.Get(consts.CodeBadRequest))
			},
			inPath: "/api/v1/product/1000",
			inParams: map[string]string {
				"id": "1",
			},
			inProduct: &models.Product {
				Title: "WB",
				Price: -1000,
			},
			expStatusCode: errors.Get(consts.CodeBadRequest).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeBadRequest),
			},
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64, prod *models.Product) {
				productUse.
					EXPECT().
					UpdateById(prodId, prod).
					Return(errors.Get(consts.CodeInternalError))
			},
			inPath: "/api/v1/product/1000",
			inParams: map[string]string {
				"id": "1",
			},
			inProduct: &models.Product {
				Title: "WB",
				Price: 500,
			},
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("POST", testCase.inPath, converter.AnyBytesToString(testCase.inProduct))
			r  = mux.SetURLVars(r, testCase.inParams)
			w := httptest.NewRecorder()
			productUse := mock_product.NewMockProductUsecase(ctrl)

			productId, _ := mux.Vars(r)["id"]
			intProductId, _ := strconv.ParseUint(productId, 10, 64)
			testCase.mockBehaviour(productUse, intProductId, testCase.inProduct)
			productHandler := NewProductHandler(productUse)
			productHandler.Configure(mx, nil)

			productHandler.UpdateProductById()(w, r)
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}

func TestProductHandler_DeleteProductById(t *testing.T) {
	type mockBehaviour func(productUse *mock_product.MockProductUsecase, prodId uint64)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inPath string
		inParams map[string]string
		inProduct *models.Product
		expStatusCode int
		expRespBody response.Response
	}{
		{
			name: "OK",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64) {
				productUse.
					EXPECT().
					DeleteById(prodId).
					Return(nil)
			},
			inPath: "/api/v1/product/1",
			inParams: map[string]string {
				"id": "1",
			},
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"deleted_elements": true,
				},
			},
		},
		{
			name: "Error: CodeProductDoesNotExist",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64) {
				productUse.
					EXPECT().
					DeleteById(prodId).
					Return(errors.Get(consts.CodeProductDoesNotExist))
			},
			inPath: "/api/v1/product/1000",
			inParams: map[string]string {
				"id": "1000",
			},
			expStatusCode: errors.Get(consts.CodeProductDoesNotExist).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeProductDoesNotExist),
			},
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(productUse *mock_product.MockProductUsecase, prodId uint64) {
				productUse.
					EXPECT().
					DeleteById(prodId).
					Return(errors.Get(consts.CodeInternalError))
			},
			inPath: "/api/v1/product/1",
			inParams: map[string]string {
				"id": "1",
			},
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("DELETE", testCase.inPath, converter.AnyBytesToString(testCase.inProduct))
			r  = mux.SetURLVars(r, testCase.inParams)
			w := httptest.NewRecorder()
			productUse := mock_product.NewMockProductUsecase(ctrl)

			productId, _ := mux.Vars(r)["id"]
			intProductId, _ := strconv.ParseUint(productId, 10, 64)
			testCase.mockBehaviour(productUse, intProductId)
			productHandler := NewProductHandler(productUse)
			productHandler.Configure(mx, nil)

			productHandler.DeleteProductById()(w, r)
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}

