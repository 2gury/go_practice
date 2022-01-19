package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gorilla/mux"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var html = []byte(`
<html>
	<body>
	<form action="/upload" method="post" enctype="multipart/form-data">
		Image: <input type="file" name="my_file">
		<input type="submit" value="Upload">
	</form>
	</body>
</html>
`)

func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write(html)
}


func UploadPage(w http.ResponseWriter, r *http.Request) {
	data, _, err := r.FormFile("my_file")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer data.Close()

	randName := RandStringRunes(10)
	tmpFile := "./images/" + randName + ".jpg"
	newFile, err := os.Create(tmpFile)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	hasher := md5.New()
	bytes, err := io.Copy(newFile, io.TeeReader(data, hasher))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	newFile.Sync()
	newFile.Close()

	md5Sum := hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", MainPage)
	mux.HandleFunc("/upload", UploadPage)
	log.Printf("starting server at port :8080")
	http.ListenAndServe(":8080", mux)
}
