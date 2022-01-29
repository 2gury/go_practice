package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

var messagesTmpl = `<html>
<body>
	&lt;form action=&quot;http://jsonplaceholder.typicode.com/posts&quot; method=&quot;POST&quot;&gt;
	<br />
	&lt;input type=&quot;submit&quot; value=&quot;Click me&quot;&gt;
	<br />
	&lt;/form&gt;
	<br />

	<form action="/comment" method="post">
		<textarea name="comment"></textarea><br />
		<input type="submit" value="Comment">
	</form>
	<br />
	
    {{range $idx, $var := .Messages}}
		<div style="border: 1px solid black; padding: 5px; margin: 5px;">
			{{$var.Message}}
		</div>
    {{end}}
</body></html>`

type Msg struct {
	ID int
	Message string
	Rating int
}

var messages = map[int]*Msg{}
var lastId = 0

func Root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("main")
	tmpl, _ = tmpl.Parse(messagesTmpl)
	tmpl.Execute(w, struct {
		Messages map[int]*Msg
	}{
		Messages: messages,
	})
}

func Comment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dataComment := r.FormValue("comment")
	idComment := lastId
	messages[lastId] = &Msg{
		ID: idComment,
		Message: dataComment,
		Rating: 0,
	}
	lastId++
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Root).Methods("GET")
	mux.HandleFunc("/comment", Comment).Methods("POST")
	log.Println("Launch server at :8080 port")
	http.ListenAndServe(":8080", mux)
}
