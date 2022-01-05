package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
)

type SendMessage struct {
	Id int `valid:",optional"`
	Name string `schema:"from" valid:"email"`
	Message string `valid:"message,required"`
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Request " + r.URL.String() + "\n\n"))

	msg := &SendMessage{}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err := decoder.Decode(msg, r.URL.Query())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Msg: %#v\n\n", msg)))

	_, err = govalidator.ValidateStruct(msg)
	if err != nil {
		if allErrs, ok := err.(govalidator.Errors); ok {
			for _, fld := range allErrs.Errors() {
				data := []byte(fmt.Sprintf("field: %#v\n\n", fld))
				w.Write(data)
			}
		}
		w.Write([]byte(fmt.Sprintf("error: %s\n\n", err)))
	} else {
		w.Write([]byte(fmt.Sprintf("message is correct")))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootPage)
	http.ListenAndServe(":8080", r)
}

func init() {
	govalidator.CustomTypeTagMap.Set(
		"message",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			if len(subject) > 10 {
				return false
			}
			return true
		}),
		)
}
