package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var secret = []byte("somethingSecretPhrase")

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": r.FormValue("username"),
		})

		str, err := token.SignedString(secret)
		if err != nil {
			http.Error(w, "{Err when set token}", http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name: "username",
			Value: str,
			Expires: time.Now().Add(time.Hour),
		}

		http.SetCookie(w, &cookie)
		w.Write([]byte(str))
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("username")
		if err == http.ErrNoCookie {
			w.Write([]byte("You are not logged"))
			return
		}
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			w.Write([]byte("hello " + claims["username"].(string)))
			return
		}
		w.Write([]byte("You are not logged"))
		fmt.Println(err)
	})
	http.ListenAndServe(":8080", r)
}