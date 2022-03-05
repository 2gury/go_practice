package main

import (
	"context"
	"go_practice/18_protobuf_stream/summator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"sync"
	"time"
)

func main() {
	grpcConn, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcConn.Close()

	summ := summator.NewSummatorClient(grpcConn)
	stream, err := summ.Sum(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		nums := []*summator.Nums {
			&summator.Nums{
				FirstNum: 1,
				SecondNum: 2,
			},
			&summator.Nums{
				FirstNum: 5,
				SecondNum: 7,
			},
			&summator.Nums{
				FirstNum: 8,
				SecondNum: 3,
			},
		}
		for _, num := range nums {
			time.Sleep(time.Second)
			log.Printf("-> %v\n", num)
			stream.Send(num)
		}
		stream.CloseSend()
		log.Println("Close sending")
	} (wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("stream closed")
				return
			}
			if err != nil {
				log.Println("error stream ", err)
				return
			}
			log.Printf("<- %v\n", res)
		}
	} (wg)
	wg.Wait()
}
