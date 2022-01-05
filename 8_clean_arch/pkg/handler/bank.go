package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) GetBanks(w http.ResponseWriter, r *http.Request) {
	banks, _ := h.Service.GetBanks()
	w.Write([]byte(fmt.Sprintf("%#v", banks)))
}

func (h *Handler) GetBankById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bank by id"))
}

func (h *Handler) GetBankByCity(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bank by city"))
}
