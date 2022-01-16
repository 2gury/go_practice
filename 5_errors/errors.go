package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type ResourceError struct {
	URL  string
	Err  error
	Code int
}

func (e *ResourceError) Error() string {
	return fmt.Sprintf("Err: %s at URL: %s with Code: %d", e.Err, e.URL, e.Code)
}

var (
	client = http.Client{Timeout: time.Duration(time.Second)}

	ErrNoResource = errors.New("No resource")
)

func getRemoteResourceBaseError() error {
	url := "http://127.0.0.1:9999/pages?id=123"
	_, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("getRemoteResource: %s at %s", err, url)
	}
	return nil
}

func getRemoteResourceNamedError() error {
	return ErrNoResource
}

func getRemoteResourceWrapedError() error {
	url := "http://127.0.0.1:9999/pages?id=123"
	_, err := client.Get(url)
	if err != nil {
		return errors.Wrap(err, "No resource")
	}
	return nil
}

func getRemoteResourceOwnError() error {
	url := "http://127.0.0.1:9999/pages?id=123"
	_, err := client.Get(url)
	if err != nil {
		return &ResourceError{
			URL:  url,
			Err:  err,
			Code: http.StatusInternalServerError,
		}
	}
	return nil
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	typeOfError := r.FormValue("error")
	var err error
	switch typeOfError {
	case "base":
		err = getRemoteResourceBaseError()
	case "named":
		err = getRemoteResourceNamedError()
	case "wraped":
		err = errors.Cause(getRemoteResourceWrapedError())
	case "own":
		err = errors.Cause(getRemoteResourceOwnError())
	}
	if err != nil {
		fmt.Printf("error happened: %+v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", errorPage)
	http.ListenAndServe(":8080", r)
}
