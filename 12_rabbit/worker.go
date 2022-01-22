package main

import (
	"encoding/json"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
	"image/jpeg"
	"log"
	"os"
	"sync"
	"time"
)

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

	_, err = rabbitChan.QueueDeclare(
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

	err = rabbitChan.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	tasks, err := rabbitChan.Consume(
		ImageResizeQueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go ResizeWorker(tasks)

	log.Print("\nworker started")
	wg.Wait()
}

func ResizeWorker(tasks <- chan amqp.Delivery) {
	for taskItem := range tasks {
		log.Printf("\nInput task: %+v", taskItem)

		task := &ImgResizeTask{}
		err := json.Unmarshal(taskItem.Body, task)
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		originalPath := fmt.Sprintf("./images/%s.jpg", task.MD5)
		for _, size := range sizes {
			resizedPath := fmt.Sprintf("./images/%s_%d.jpg", task.MD5, size)
			err := ResizeImage(originalPath, resizedPath, size)
			if err != nil {
				log.Printf(err.Error())
				continue
			}
			time.Sleep(5 * time.Second)
		}

		taskItem.Ack(false)
	}
}

func ResizeImage(originalPath string, resizedPath string, size uint) error {
	file, err := os.Open(originalPath)
	if err != nil {
		return fmt.Errorf("cant open file %s: %s", originalPath, err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		return fmt.Errorf("cant jpeg decode file %s", err)
	}
	file.Close()

	resizeImage := resize.Resize(size, 0, img, resize.Lanczos3)

	out, err := os.Create(resizedPath)
	if err != nil {
		return fmt.Errorf("cant create file %s: %s", resizedPath, err)
	}
	defer out.Close()

	jpeg.Encode(out, resizeImage, nil)

	return nil
}
