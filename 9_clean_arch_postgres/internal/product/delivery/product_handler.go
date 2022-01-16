package delivery

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
	"go_practice/8_clean_arch/tools/response"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productUse product.ProductUsecase
}

func NewProductHandler(use product.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUse: use,
	}
}

func (h *ProductHandler) Configure(m *mux.Router) {
	m.HandleFunc("/product", h.GetProducts).Methods("GET")
	m.HandleFunc("/product/", h.AddProduct).Methods("PUT")
	m.HandleFunc("/product/{id:[0-9]+}", h.UpdateProductById).Methods("POST")
	m.HandleFunc("/product/{id:[0-9]+}", h.DeleteProductById).Methods("DELETE")
	m.HandleFunc("/product/{id:[0-9]+}", h.GetProductById).Methods("GET")
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productUse.List()
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error when get products}",
		})
		return
	}

	json.NewEncoder(w).Encode(response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"products": products,
		},
	})
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId, _ := mux.Vars(r)["id"]
	ok := govalidator.IsInt(productId)
	if !ok {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error can't validate input params}",
		})
		return
	}
	intProductId, _ := strconv.Atoi(productId)
	product, err := h.productUse.GetById(uint64(intProductId))
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error when get product by id}",
		})
		return
	}
	if product == nil {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{No product with this id}",
		})
		return
	}

	json.NewEncoder(w).Encode(response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"product": product,
		},
	})
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	ok, _ := govalidator.ValidateStruct(product)
	if !ok {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error can't validate input params}",
		})
		return
	}
	id, err := h.productUse.Create(product)
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error when add product}",
		})
		return
	}
	json.NewEncoder(w).Encode(response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"id": id,
		},
	})
}

func (h *ProductHandler) UpdateProductById(w http.ResponseWriter, r *http.Request) {
	productId, _ := mux.Vars(r)["id"]
	ok := govalidator.IsInt(productId)
	if !ok {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error can't validate input params}",
		})
		return
	}
	intProductId, _ := strconv.Atoi(productId)
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	ok, _ = govalidator.ValidateStruct(product)
	if !ok {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error can't validate input params}",
		})
		return
	}
	updated, err := h.productUse.UpdateById(uint64(intProductId), product)
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error when update product}",
		})
		return
	}
	if !updated {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error no product found}",
		})
		return
	}
	json.NewEncoder(w).Encode(response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"update": updated,
		},
	})
}

func (h *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	productId, _ := mux.Vars(r)["id"]
	ok := govalidator.IsInt(productId)
	if !ok {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error can't validate input params}",
		})
		return
	}
	intProductId, _ := strconv.Atoi(productId)
	deleted, err := h.productUse.DeleteById(uint64(intProductId))
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error when delete product by id}",
		})
		return
	}
	if !deleted {
		json.NewEncoder(w).Encode(response.Response{
			Code:  http.StatusBadRequest,
			Error: "{Error no product found}",
		})
		return
	}
	json.NewEncoder(w).Encode(response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"deleted_elements": deleted,
		},
	})
}
