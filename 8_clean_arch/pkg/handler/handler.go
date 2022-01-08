package handler

import (
	"github.com/gorilla/mux"
	"go_practice/8_clean_arch/pkg/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		Service: svc,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	handler := mux.NewRouter()

	handler.HandleFunc("/product", h.GetProducts).Methods("GET")
	handler.HandleFunc("/product/", h.AddProduct).Methods("POST")
	handler.HandleFunc("/product/{id:[0-9]+}", h.GetProductById).Methods("GET")

	return handler
}
