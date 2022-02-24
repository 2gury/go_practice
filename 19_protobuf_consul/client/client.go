package main

import (
	"context"
	consulapi "github.com/hashicorp/consul/api"
	registry "github.com/liyue201/grpc-lb/registry/consul"
	"go_practice/19_protobuf_consul/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"time"
)

var consulAddr = "127.0.0.1:8500"

func main() {
	registry.RegisterResolver("consul", &consulapi.Config{Address: consulAddr}, "echo-api")

	grpcConn, err := grpc.Dial(
		"consul:///",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcConn.Close()

	echoSevice := echo.NewEchoClient(grpcConn)
	step := 0
	for {
		output, err := echoSevice.Say(context.Background(), &echo.Input{Message: strconv.Itoa(step)})
		log.Printf("get output: %s %v", output.Message, err)
		time.Sleep(2 * time.Second)
		step++
	}
}

