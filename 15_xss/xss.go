package main

import (
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
	"log"
	"net/http"
)

var messages = []string{"Lulz", "keke"}

var sanitizer = bluemonday.UGCPolicy()

var messagesTmpl = `<html><body>
	&lt;script&gt;alert(1)&lt;/script&gt;
	<br />
	<br />
	<form action="/comment" method="post">
		<textarea name="comment"></textarea><br />
		<input type="submit" value="Comment">
	</form>
    {{range .Messages}}
		<div style="border: 1px solid black; padding: 5px; margin: 5px;">
			<!-- text/template по-умолч ничего не экранируется, надо указать | html --> 
			<!-- html/template по-умолч будет экранировать --> 

			{{.}}
		</div>
    {{end}}
</body></html>`

func Root(w http.ResponseWriter, r *http.Request) {
	param := r.FormValue("param")
	outputMessages := []template.HTML{}
	tmpl := template.New("main")
	
	tmpl, _ = tmpl.Parse(messagesTmpl)
	if param == "sanitize" {
		for _, v := range messages {
			outputMessages = append(outputMessages, template.HTML(sanitizer.Sanitize(v)))
		}
	} else {
		for _, v := range messages {
			outputMessages = append(outputMessages, template.HTML(v))
		}
	}
	tmpl.Execute(w, struct {
		Messages []template.HTML
	}{
		Messages: outputMessages,
	})
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	commentText := r.FormValue("comment")
	
	messages = append(messages, commentText)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Root)
	mux.HandleFunc("/comment", AddComment)
	log.Println("Launch server at :8080 port")
	http.ListenAndServe(":8080", mux)
}
