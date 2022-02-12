package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/mwares"
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

func (h *ProductHandler) Configure(m *mux.Router, mwManager *mwares.MiddlewareManager) {
	m.HandleFunc("/api/v1/product", h.GetProducts()).Methods("GET", "OPTIONS")
	m.HandleFunc("/api/v1/product/", h.AddProduct()).Methods("PUT", "OPTIONS")
	m.HandleFunc("/api/v1/product/{id:[0-9]+}", h.UpdateProductById()).Methods("POST", "OPTIONS")
	m.HandleFunc("/api/v1/product/{id:[0-9]+}", h.DeleteProductById()).Methods("DELETE", "OPTIONS")
	m.HandleFunc("/api/v1/product/{id:[0-9]+}", h.GetProductById()).Methods("GET", "OPTIONS")
}

func (h *ProductHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		products, err := h.productUse.List()
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}

		contextHelper.WriteStatusCodeContext(ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
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
		ctx := r.Context()
		productId, _ := mux.Vars(r)["id"]
		intProductId, parseErr := strconv.ParseUint(productId, 10, 64)

		if parseErr != nil {
			err := errors.Get(consts.CodeValidateError)
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		product, err := h.productUse.GetById(intProductId)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
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
		defer r.Body.Close()
		ctx := r.Context()
		req := &Request{}

		json.NewDecoder(r.Body).Decode(&req)
		if err := request_reader.ValidateStruct(req); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		product := models.Product{
			Title: req.Title,
			Price: req.Price,
		}
		id, err := h.productUse.Create(product)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
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
		defer r.Body.Close()
		ctx := r.Context()
		productId, _ := mux.Vars(r)["id"]
		intProductId, parseErr := strconv.ParseUint(productId, 10, 64)

		if parseErr != nil {
			err := errors.Get(consts.CodeValidateError)
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		req := &Request{}
		json.NewDecoder(r.Body).Decode(&req)
		if err := request_reader.ValidateStruct(req); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		product := models.Product{
			Title: req.Title,
			Price: req.Price,
		}
		if err := h.productUse.UpdateById(intProductId, product); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
			Body: &response.Body{
				"updated_elements": true,
			},
		})
	}
}

func (h *ProductHandler) DeleteProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		productId, _ := mux.Vars(r)["id"]
		intProductId, parseErr := strconv.ParseUint(productId, 10, 64)

		if parseErr != nil {
			err := errors.Get(consts.CodeValidateError)
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		if err := h.productUse.DeleteById(intProductId); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
			Body: &response.Body{
				"deleted_elements": true,
			},
		})
	}
}
