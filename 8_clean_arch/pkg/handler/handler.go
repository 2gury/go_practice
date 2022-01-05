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

	handler.HandleFunc("/bank", h.GetBanks).Methods("GET")
	handler.HandleFunc("/bank/{id:[0-9]+}", h.GetBankById).Methods("GET")
	handler.HandleFunc("/bank/{cityName}", h.GetBankByCity).Methods("GET")

	handler.HandleFunc("/city", h.GetCities).Methods("GET")
	handler.HandleFunc("/city/{id:[0-9]+}", h.GetCityById).Methods("GET")

	return handler
}
