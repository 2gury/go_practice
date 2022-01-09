package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	store *BookStore
}

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.store.GetBooks()
	if err != nil {
		http.Error(w, "{Error while get books}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{}{
		"body": books,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *Handler) AddBook(w http.ResponseWriter, r *http.Request) {
	var bookInput BookInput
	json.NewDecoder(r.Body).Decode(&bookInput)
	id, err := h.store.AddBook(bookInput)
	if err != nil {
		http.Error(w, "{Error while add book}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{}{
		"body": id,
	}
	json.NewEncoder(w).Encode(body)
}

func (h *Handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "{Error while get book by id}", http.StatusInternalServerError)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "{Error while get book by id}", http.StatusInternalServerError)
		return
	}
	book, err := h.store.GetBookById(idInt)
	if err != nil {
		http.Error(w, "{Error while get book by id}", http.StatusInternalServerError)
		return
	}
	if book == nil {
		http.Error(w, "{No book with this id}", http.StatusInternalServerError)
		return
	}
	body := map[string]interface{}{
		"body": book,
	}
	json.NewEncoder(w).Encode(body)
}

func main() {
	handler := Handler{NewBookStore()}

	r := mux.NewRouter()
	r.HandleFunc("/", handler.GetBooks).Methods("GET")
	r.HandleFunc("/book/", handler.AddBook).Methods("PUT")
	r.HandleFunc("/book/{id:[0-9]+}", handler.GetBookById).Methods("GET")
	http.ListenAndServe(":8080", r)
}
