package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var loginFormTmpl = `
<html>
	<body>
	<form action="/login" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`

type Handler struct {
	db *sql.DB
}

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(loginFormTmpl))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var id int
	var login string

	loginInput := r.FormValue("login")
	query := fmt.Sprintf("SELECT id, login FROM users WHERE login = '%s' LIMIT 1", loginInput)

	w.Write([]byte(query))
	err := h.db.QueryRow(query).Scan(&id, &login)
	if err == sql.ErrNoRows {
		w.Write([]byte("\nError"))
	} else {
		w.Write([]byte(fmt.Sprintf("\nid = %d, login = %s", id, login)))
	}

	err = h.db.QueryRow(`SELECT id, login FROM users WHERE login = ? LIMIT 1`, loginInput).
		Scan(&id, &login)
	if err == sql.ErrNoRows {
		w.Write([]byte("\nError"))
	} else {
		w.Write([]byte(fmt.Sprintf("\nid = %d, login = %s", id, login)))
	}
}

func main() {
	connString := "root:love@tcp(localhost:3307)/golang?&charset=utf8&interpolateParams=true"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)

	h := Handler{
		db: db,
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/", h.Root)
	mux.HandleFunc("/login", h.Login)
	log.Printf("start server at :8080")
	http.ListenAndServe(":8080", mux)
}
