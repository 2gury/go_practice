package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	redisAddr = "redis://user:@localhost:6379/0"

	users = map[string]string {
		"user": "zurg",
	}

	sessKeyLen = 10
	sessManager *SessionManager

	loginFormTmpl = []byte(`
	<html>
		<body>
		<form action="/login" method="post">
			Login: <input type="text" name="login">
			Password: <input type="password" name="password">
			<input type="submit" value="Login">
		</form>
		</body>
	</html>
	`)
)

func CheckSession(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		return nil, err
	}
	sess, err := sessManager.Check(&SessionId{
		Id: cookie.Value,
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func RootPage(w http.ResponseWriter, r *http.Request) {
	sess, err := CheckSession(r)
	if err != nil {
		if err == http.ErrNoCookie {
			w.Write([]byte(loginFormTmpl))
			return
		}
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", sess.Login)))
	w.Write([]byte(`<a href="/logout">logout</a>`))
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	inputLogin := r.FormValue("login")
	inputPassword := r.FormValue("password")

	pass, ok := users[inputLogin]
	if !ok || pass != inputPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id, err := sessManager.Create(&Session{
		Login: inputLogin,
		Useragent: r.UserAgent(),
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	cookie := &http.Cookie{
		Name: "user_id",
		Value: id.Id,
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sessManager.Delete(&SessionId{
		Id: cookie.Value,
	})
	cookie.Expires = time.Now().Add(-24 *time.Hour)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	redisConn, err := redis.DialURL(redisAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	sessManager = NewSessionManager(redisConn)

	mux := mux.NewRouter()
	mux.HandleFunc("/", RootPage)
	mux.HandleFunc("/login", LoginPage)
	mux.HandleFunc("/logout", LogoutPage)
	log.Printf("\nstarted server at 8080 port")
	http.ListenAndServe(":8080", mux)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
