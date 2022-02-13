package delivery

import (
	"github.com/golang/mock/gomock"
	mux "github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/models"
	mock_product "go_practice/9_clean_arch_db/internal/product/mocks"
	"go_practice/9_clean_arch_db/tools/converter"
	"go_practice/9_clean_arch_db/tools/response"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestProductHandler_GetProducts(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productUse := mock_product.NewMockProductUsecase(ctrl)

	products := []*models.Product{
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

	r := httptest.NewRequest("GET", "/api/v1/product", nil)
	w := httptest.NewRecorder()

	productHandler := NewProductHandler(productUse)
	mux := mux.NewRouter()
	productHandler.Configure(mux, nil)

	productUse.
		EXPECT().
		List().
		Return(products, nil)

	resp := response.Response{
		Body: &response.Body{
			"products": products,
		},
	}

	productHandler.GetProducts()(w, r)

	expResBody, err := converter.AnyToBytesBuffer(resp)
	if err != nil {
		t.Error(err.Error())
		return
	}
	bytes, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, expResBody.String(), string(bytes))
}
