package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/8_clean_arch/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.GetProducts()
	if err != nil {
		http.Error(w, "{Error while get products}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{} {
		"body": products,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
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
	product, err := h.Service.GetProductById(intProductId)
	if err != nil {
		http.Error(w, "{Error while get product by id}", http.StatusInternalServerError)
		return
	}
	if product == nil {
		http.Error(w, "{No product with this id}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{} {
		"body": product,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	id, err := h.Service.AddProduct(product)
	if err != nil {
		http.Error(w, "{Error while add product}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{} {
		"id": id,
	}
	json.NewEncoder(w).Encode(body)
}
