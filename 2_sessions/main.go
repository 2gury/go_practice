package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Handler struct {
	users map[string]User
	sessions map[string]int
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var userInput UserInput
	json.NewDecoder(r.Body).Decode(&userInput)

	user, ok := h.users[userInput.Username]
	if !ok {
		http.Error(w, "{No user}", http.StatusInternalServerError)
		return
	}
	if user.Password != userInput.Password {
		http.Error(w, "{Err password}", http.StatusInternalServerError)
		return
	}
	SID := RandStringRunes(32)
	cookie := http.Cookie{
		Name: "user_id",
		Value: SID,
		Expires: time.Now().Add(time.Hour),
	}
	h.sessions[SID] = user.Id
	http.SetCookie(w, &cookie)
	w.Write([]byte(SID))
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err == http.ErrNoCookie {
		return
	}
	_, ok := h.sessions[cookie.Value]
	if !ok {
		http.Error(w, "{Err cookie}", http.StatusInternalServerError)
		return
	}
	delete(h.sessions, cookie.Value)
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
}

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err == http.ErrNoCookie {
		w.Write([]byte("No logged in"))
		return
	}
	_, ok := h.sessions[cookie.Value]
	if !ok {
		w.Write([]byte("No logged in"))
		return
	}
	w.Write([]byte("You are logged in"))
}

func main() {
	h := Handler{
		users: map[string]User{
			"logog": {
				Id: 0,
				Username: "logog",
				Password: "123456",
			},
		},
		sessions: map[string]int{},
	}
	r := mux.NewRouter()

	r.HandleFunc("/", h.Root)
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/logout", h.Logout)

	http.ListenAndServe(":8080", r)
}