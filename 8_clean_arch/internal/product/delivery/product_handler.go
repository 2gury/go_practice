package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
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
	m.HandleFunc("/product/{id:[0-9]+}", h.UpdateProduct).Methods("POST")
	m.HandleFunc("/product/{id:[0-9]+}", h.GetProductById).Methods("GET")
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.u.List()
	if err != nil {
		http.Error(w, "{Error while get product}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{}{
		"body": products,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "{Error when get product id}", http.StatusInternalServerError)
		return
	}
	intProductId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "{Error while get product by id}", http.StatusInternalServerError)
		return
	}
	product, err := h.u.GetById(uint64(intProductId))
	if err != nil {
		http.Error(w, "{Error while get product by id}", http.StatusInternalServerError)
		return
	}
	if product == nil {
		http.Error(w, "{No product with this id}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{}{
		"body": product,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.ProductInput
	json.NewDecoder(r.Body).Decode(&product)
	id, err := h.u.Create(product)
	if err != nil {
		http.Error(w, "{Error while add product}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{}{
		"id": id,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "{Error when get product id}", http.StatusInternalServerError)
		return
	}
	intProductId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "{Error while get product by id}", http.StatusInternalServerError)
		return
	}
	var product models.ProductInput
	json.NewDecoder(r.Body).Decode(&product)
	numUpdated, err := h.u.UpdateById(uint64(intProductId), product)
	if err != nil {
		http.Error(w, "{Error while update product}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{}{
		"updated_elements": numUpdated,
	}
	json.NewEncoder(w).Encode(body)
}

