package handler

import "net/http"

func GetBanks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("banks"))
}

func GetBankById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bank by id"))
}

func GetBankByCity(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bank by city"))
}
