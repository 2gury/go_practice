package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
	"go_practice/8_clean_arch/tools/response"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	u product.ProductUsecase
}

func NewProductHandler(use product.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		u: use,
	}
}

func (h *ProductHandler) Configure(m *mux.Router) {
	m.HandleFunc("/product", h.GetProducts).Methods("GET")
	m.HandleFunc("/product/", h.AddProduct).Methods("POST")
	m.HandleFunc("/product/{id:[0-9]+}", h.UpdateProductById).Methods("POST")
	m.HandleFunc("/product/{id:[0-9]+}", h.DeleteProductById).Methods("DELETE")
	m.HandleFunc("/product/{id:[0-9]+}", h.GetProductById).Methods("GET")
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.u.List()
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusOK,
			Error: "{Error while get product}",
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
	productId, ok := mux.Vars(r)["id"]
	if !ok {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error when get product id}",
		})
		return
	}
	intProductId, err := strconv.Atoi(productId)
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error while get product by id}",
		})
		return
	}
	product, err := h.u.GetById(uint64(intProductId))
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error while get product by id}",
		})
		return
	}
	if product == nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
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
	id, err := h.u.Create(product)
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error while add product}",
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
	productId, ok := mux.Vars(r)["id"]
	if !ok {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error when get product id}",
		})
		return
	}
	intProductId, err := strconv.Atoi(productId)
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error while get product by id}",
		})
		return
	}
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	numUpdated, err := h.u.UpdateById(uint64(intProductId), product)
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error while update product}",
		})
		return
	}
	json.NewEncoder(w).Encode(response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"updated_elements": numUpdated,
		},
	})
}

func (h *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	productId, ok := mux.Vars(r)["id"]
	if !ok {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error when get product id}",
		})
		return
	}
	intProductId, err := strconv.Atoi(productId)
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error while get product by id}",
		})
		return
	}
	numDeleted, err := h.u.DeleteById(uint64(intProductId))
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error while delete product by id}",
		})
		return
	}
	json.NewEncoder(w).Encode(response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"deleted_elements": numDeleted,
		},
	})
}

