package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"text/template"
	"log"
	"net/http"
	"strconv"
)

var messagesTmpl = `<html>
<head>
<script>
	function rateComment(id, vote) {
		var request = new XMLHttpRequest();
		request.open('POST', '/rate?id='+id+"&vote="+(vote > 0 ? "up" : "down"), true);

		request.onload = function() {
		    var resp = JSON.parse(request.responseText);
			console.log(resp, resp.id)
			console.log('#rating-'+resp.id)
			console.log(document.querySelector('#rating-'+resp.id))
			document.querySelector('#rating-'+resp.id).innerHTML = resp.rating;
		};
		request.send();
	}
</script>
</head>
<body>

	&lt;form action=&quot;/rate?id=0&amp;vote=up&quot; method=&quot;POST&quot;&gt;
	<br />
	&lt;input type=&quot;hidden&quot; name=&quot;id&quot; value=&quot;0&quot;&gt;
	<br />
	&lt;input type=&quot;hidden&quot; name=&quot;vote&quot; value=&quot;down&quot;&gt;
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
			<button onclick="rateComment({{$var.ID}}, 1)">&uarr;</button>
			<span id="rating-{{$var.ID}}">{{$var.Rating}}</span>
			<button onclick="rateComment({{$var.ID}}, -1)">&darr;</button>
			&nbsp;
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

func Rate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	emptyResponse := []byte(`{"id":0, "rating":0}`)
	idComment := r.URL.Query().Get("id")
	voteComment := r.URL.Query().Get("vote")
	var rateComment int
	if voteComment == "up" {
		rateComment = 1
	} else if voteComment == "down" {
		rateComment = -1
	} else {
		w.Write(emptyResponse)
		return
	}
	intIdComment, err := strconv.Atoi(idComment)
	if err != nil {
		w.Write(emptyResponse)
		return
	}
	message, ok := messages[intIdComment]
	if !ok {
		w.Write(emptyResponse)
		return
	}
	message.Rating += rateComment
	w.Write([]byte(fmt.Sprintf(`{"id":%d, "rating":%d}`, message.ID, message.Rating)))
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Root).Methods("GET")
	mux.HandleFunc("/comment", Comment).Methods("POST")
	mux.HandleFunc("/rate", Rate).Methods("POST")
	log.Println("Launch server at :8080 port")
	http.ListenAndServe(":8080", mux)
}
