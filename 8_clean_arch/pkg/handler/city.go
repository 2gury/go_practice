package handler

import "net/http"

func (h *Handler) GetCities(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cities"))
}

func (h *Handler) GetCityById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("city by id"))
}
