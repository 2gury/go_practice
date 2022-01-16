package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("Hello"))
	} else {
		w.Write([]byte("Hello, " + cookie.Value))
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "username",
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return
	}
	cookie.Expires = time.Now().Add(-24 * time.Hour)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func adminPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="/">site index</a>`))
	w.Write([]byte(`Admin page`))
}

func panicPage(w http.ResponseWriter, r *http.Request) {
	panic("Panic must be covered")
}

func adminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("adminAuthMiddleware")
		_, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println("Access denied")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func panicCoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			fmt.Println("panicCoverMiddleware")
			if err := recover(); err != nil {
				fmt.Println("Panic recovered", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accessLogMiddleware")
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s %s %s\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func main() {
	adminMux := mux.NewRouter()
	adminMux.HandleFunc("/admin", adminPage)
	adminHandler := adminAuthMiddleware(adminMux)

	mux := mux.NewRouter()
	mux.Handle("/admin", adminHandler)
	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/login", loginPage)
	mux.HandleFunc("/panic", panicPage)
	mux.HandleFunc("/logout", logoutPage)

	handler := accessLogMiddleware(mux)
	handler = panicCoverMiddleware(handler)

	http.ListenAndServe(":8080", handler)
}
