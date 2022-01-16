package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_practice/11_clean_arch_mongo_db/internal"
	"go_practice/11_clean_arch_mongo_db/internal/models"
	"go_practice/11_clean_arch_mongo_db/tools/response"
	"net/http"
)

type ProductHandler struct {
	u internal.ProductUsecase
}

func NewProductHandler(use internal.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		u: use,
	}
}

func (h *ProductHandler) Configure(mux *mux.Router) {
	mux.HandleFunc("/product", h.List).Methods("GET")
	mux.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	mux.HandleFunc("/product/", h.AddProduct).Methods("POST")
	mux.HandleFunc("/product/{id}", h.UpdateProduct).Methods("POST")
	mux.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.u.GetAllProducts()
	if err != nil {
		json.NewEncoder(w).Encode(&response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error when get products}",
		})
		return
	}
	json.NewEncoder(w).Encode(&response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"body": products,
		},
	})
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		json.NewEncoder(w).Encode(&response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error format id}",
		})
		return
	}
	product, err := h.u.GetProductById(id)
	if err != nil {
		json.NewEncoder(w).Encode(&response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error when get product by id}",
		})
		return
	}
	json.NewEncoder(w).Encode(&response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"body": product,
		},
	})
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	json.NewDecoder(r.Body).Decode(product)
	defer r.Body.Close()

	lastId, err := h.u.AddProduct(product)
	if err != nil {
		json.NewEncoder(w).Encode(&response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error when add product}",
		})
		return
	}
	json.NewEncoder(w).Encode(&response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"id": lastId,
		},
	})
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		json.NewEncoder(w).Encode(&response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error format id}",
		})
		return
	}
	product := &models.Product{}
	json.NewDecoder(r.Body).Decode(product)
	defer r.Body.Close()
	product.Id = hexId
	affected, err := h.u.UpdateProduct(product)
	if err != nil {
		json.NewEncoder(w).Encode(&response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error when update product}",
		})
		return
	}
	json.NewEncoder(w).Encode(&response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"affected": affected,
		},
	})
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	deleted, err := h.u.DeleteProductById(id)
	if err != nil {
		json.NewEncoder(w).Encode(&response.Response{
			Code: http.StatusBadRequest,
			Error: "{Error when delete product}",
		})
		return
	}
	json.NewEncoder(w).Encode(&response.Response{
		Code: http.StatusOK,
		Body: &response.Body{
			"deleted": deleted,
		},
	})
}