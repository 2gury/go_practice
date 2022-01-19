package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/product"
	"go_practice/9_clean_arch_db/tools/request_reader"
	"go_practice/9_clean_arch_db/tools/response"
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
	m.HandleFunc("/product", h.GetProducts()).Methods("GET")
	m.HandleFunc("/product/", h.AddProduct()).Methods("PUT")
	m.HandleFunc("/product/{id:[0-9]+}", h.UpdateProductById()).Methods("POST")
	m.HandleFunc("/product/{id:[0-9]+}", h.DeleteProductById()).Methods("DELETE")
	m.HandleFunc("/product/{id:[0-9]+}", h.GetProductById()).Methods("GET")
}

func (h *ProductHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.productUse.List()
		if err != nil {
			json.NewEncoder(w).Encode(response.Response{Code: err.HttpCode, Error: err,})
			return
		}

		json.NewEncoder(w).Encode(response.Response{Code: http.StatusOK,
			Body: &response.Body{
				"products": products,
			},
		})
	}
}

func (h *ProductHandler) GetProductById() http.HandlerFunc {
	/*
	type Query struct {
		Id int `json:"id" valid:"int,required"`
	}
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		/*
		query := &Query{}
		if err := request_reader.NewQueryReader().Read(query, r.URL.Query()); err != nil {
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		if err := request_reader.ValidateStruct(query); err != nil {
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		*/
		productId, _ := mux.Vars(r)["id"]
		intProductId, parseErr := strconv.ParseUint(productId, 10, 64)
		if parseErr != nil {
			err := errors.Get(consts.CodeValidateError)
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		product, err := h.productUse.GetById(intProductId)
		if err != nil {
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		json.NewEncoder(w).Encode(response.Response{Code: http.StatusOK,
			Body: &response.Body{
				"product": product,
			},
		})
	}
}

func (h *ProductHandler) AddProduct() http.HandlerFunc {
	type Request struct {
		Title string `json:"title" valid:"stringlength(1|64),required"`
		Price int    `json:"price" valid:"int,required"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		json.NewDecoder(r.Body).Decode(&req)
		err := request_reader.ValidateStruct(req)
		if err != nil {
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		product := models.Product{
			Title: req.Title,
			Price: req.Price,
		}
		id, err := h.productUse.Create(product)
		if err != nil {
			json.NewEncoder(w).Encode(response.Response{Code: err.HttpCode, Error: err,})
			return
		}
		json.NewEncoder(w).Encode(response.Response{ Code: http.StatusOK,
			Body: &response.Body{
				"id": id,
			},
		})
	}
}

func (h *ProductHandler) UpdateProductById() http.HandlerFunc {
	type Request struct {
		Title string `json:"title" valid:"stringlength(1|64),required"`
		Price int    `json:"price" valid:"int,required"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		productId, _ := mux.Vars(r)["id"]
		intProductId, parseErr := strconv.ParseUint(productId, 10, 64)
		if parseErr != nil {
			err := errors.Get(consts.CodeValidateError)
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		req := &Request{}
		json.NewDecoder(r.Body).Decode(&req)
		if err := request_reader.ValidateStruct(req); err != nil {
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		product := models.Product{
			Title: req.Title,
			Price: req.Price,
		}
		if err := h.productUse.UpdateById(intProductId, product); err != nil {
			json.NewEncoder(w).Encode(response.Response{Code: err.HttpCode, Error: err,})
			return
		}
		json.NewEncoder(w).Encode(response.Response{Code: http.StatusOK,
			Body: &response.Body{
				"updated_elements": true,
			},
		})
	}
}

func (h *ProductHandler) DeleteProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, _ := mux.Vars(r)["id"]
		intProductId, parseErr := strconv.ParseUint(productId, 10, 64)
		if parseErr != nil {
			err := errors.Get(consts.CodeValidateError)
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		if err := h.productUse.DeleteById(intProductId); err != nil {
			json.NewEncoder(w).Encode(response.Response{Code: err.HttpCode, Error: err,})
			return
		}
		json.NewEncoder(w).Encode(response.Response{Code: http.StatusOK,
			Body: &response.Body{
				"deleted_elements": true,
			},
		})
	}
}
