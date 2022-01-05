package handler

import "net/http"

func GetCities(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cities"))
}

func GetCityById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("city by id"))
}
