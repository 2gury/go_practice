package delivery

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"go_practice/8_clean_arch/internal/consts"
	"go_practice/8_clean_arch/internal/helpers/errors"
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
	"go_practice/8_clean_arch/tools/request_reader"
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
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
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
	return func(w http.ResponseWriter, r *http.Request) {
		productId, _ := mux.Vars(r)["id"]
		ok := govalidator.IsInt(productId)
		if !ok {
			err := errors.Get(consts.CodeValidateError)
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		intProductId, _ := strconv.Atoi(productId)
		product, err := h.productUse.GetById(uint64(intProductId))
		if err != nil {
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		if product == nil {
			err := errors.Get(consts.CodeProductDoesNotExist)
			json.NewEncoder(w).Encode(response.Response{Code: err.HttpCode, Error: err,})
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
		Title string `json:"title" valid:",required"`
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
	return func(w http.ResponseWriter, r *http.Request) {
		productId, _ := mux.Vars(r)["id"]
		ok := govalidator.IsInt(productId)
		if !ok {
			err := errors.Get(consts.CodeValidateError)
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		intProductId, _ := strconv.Atoi(productId)
		var product models.Product
		json.NewDecoder(r.Body).Decode(&product)
		err := request_reader.ValidateStruct(product)
		if err != nil {
			json.NewEncoder(w).Encode(response.Response{Code: err.HttpCode, Error: err,})
			return
		}
		err = h.productUse.UpdateById(uint64(intProductId), product)
		if err != nil {
			json.NewEncoder(w).Encode(response.Response{Code: err.HttpCode, Error: err,})
			return
		}
		json.NewEncoder(w).Encode(response.Response{Code: http.StatusOK,
			Body: &response.Body{
				"update": true,
			},
		})
	}
}

func (h *ProductHandler) DeleteProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, _ := mux.Vars(r)["id"]
		ok := govalidator.IsInt(productId)
		if !ok {
			err := errors.Get(consts.CodeValidateError)
			json.NewEncoder(w).Encode(response.Response{Code:  err.HttpCode, Error: err,})
			return
		}
		intProductId, _ := strconv.Atoi(productId)
		err := h.productUse.DeleteById(uint64(intProductId))
		if err != nil {
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
