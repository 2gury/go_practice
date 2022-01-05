package handler

import (
	"github.com/gorilla/mux"
)

type Handler struct {

}

func (h *Handler) InitRoutes() *mux.Router {
	handler := mux.NewRouter()

	handler.HandleFunc("/bank", GetBanks).Methods("GET")
	handler.HandleFunc("/bank/{id:[0-9]+}", GetBankById).Methods("GET")
	handler.HandleFunc("/bank/{cityName}", GetBankByCity).Methods("GET")

	handler.HandleFunc("/city", GetCities).Methods("GET")
	handler.HandleFunc("/city/{id:[0-9]+}", GetCityById).Methods("GET")

	return handler
}
