package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/8_clean_arch/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetBanks(w http.ResponseWriter, r *http.Request) {
	banks, err := h.Service.GetBanks()
	if err != nil {
		http.Error(w, "{Error while get banks}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{} {
		"body": banks,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *Handler) GetBankById(w http.ResponseWriter, r *http.Request) {
	bankId, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "{Error when get bank id}", http.StatusInternalServerError)
		return
	}
	intBankId, err := strconv.Atoi(bankId)
	if err != nil {
		http.Error(w, "{Error while get bank by id}", http.StatusInternalServerError)
		return
	}
	bank, err := h.Service.GetBankById(intBankId)
	if err != nil {
		http.Error(w, "{Error while get bank by id}", http.StatusInternalServerError)
		return
	}
	if bank == nil {
		http.Error(w, "{No bank with this id}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{} {
		"body": bank,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *Handler) AddBank(w http.ResponseWriter, r *http.Request) {
	var bank models.BankInput
	json.NewDecoder(r.Body).Decode(&bank)
	id, err := h.Service.AddBank(bank)
	if err != nil {
		http.Error(w, "{Error while add bank}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{} {
		"id": id,
	}
	json.NewEncoder(w).Encode(body)
}
