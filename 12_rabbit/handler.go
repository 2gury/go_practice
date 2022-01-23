package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"io"
	"log"
	"net/http"
	"os"
)


func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write(html)
}


func UploadPage(w http.ResponseWriter, r *http.Request) {
	data, handler, err := r.FormFile("my_file")
	if err != nil {
		w.Write([]byte("\nError:" + err.Error()))
		return
	}
	defer data.Close()

	randName := RandStringRunes(10)
	tmpFile := "./images/" + randName + ".jpg"
	newFile, err := os.Create(tmpFile)
	if err != nil {
		w.Write([]byte("\nError:" + err.Error()))
		return
	}

	hasher := md5.New()
	_, err = io.Copy(newFile, io.TeeReader(data, hasher))
	if err != nil {
		w.Write([]byte("\nError:" + err.Error()))
		return
	}
	newFile.Sync()
	newFile.Close()

	md5Sum := hex.EncodeToString(hasher.Sum(nil))

	realFile := "./images/" + md5Sum + ".jpg"
	if err = os.Rename(tmpFile, realFile); err != nil {
		w.Write([]byte("\nError:" + err.Error()))
		return
	}

	marshalData, _ := json.Marshal(ImgResizeTask{handler.Filename, md5Sum})
	log.Printf("Put task %s", string(marshalData))

	err = rabbitChan.Publish(
		"",
		ImageResizeQueueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body: marshalData,
		})
}

func main() {
	rabbitConn, err := amqp.Dial(rabbitAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	rabbitChan, err = rabbitConn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rabbitChan.Close()

	q, err := rabbitChan.QueueDeclare(
		ImageResizeQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("queue %s have %d msg and %d consumers\n",
		q.Name, q.Messages, q.Consumers)

	mux := mux.NewRouter()
	mux.HandleFunc("/", MainPage)
	mux.HandleFunc("/upload", UploadPage)
	log.Printf("starting server at port :8080")
	http.ListenAndServe(":8080", mux)
}
