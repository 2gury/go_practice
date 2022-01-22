package main

import (
	"github.com/streadway/amqp"
	"math/rand"
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

type ImgResizeTask struct {
	Name string
	MD5 string
}

var (
	sizes = []uint{80, 160, 320}
)

const (
	ImageResizeQueueName = "image_resize"
)

var (
	rabbitAddr = "amqp://guest:guest@127.0.0.1:5672/"
	rabbitConn *amqp.Connection
	rabbitChan *amqp.Channel
)

func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
