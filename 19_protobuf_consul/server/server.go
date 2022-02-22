package main

import (
	"flag"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"go_practice/19_protobuf_consul/echo"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

var grpcPort = flag.Int("grpc", 8081, "port for service")

func main() {
	flag.Parse()
	port := strconv.Itoa(*grpcPort)

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	echo.RegisterEchoServer(server, NewEcho(port))

	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	serviceId := "SAPI_127.0.0.1:" + port
	err = consul.Agent().ServiceRegister(&consulapi.AgentServiceRegistration{
		ID: serviceId,
		Name: "echo-api",
		Address: "127.0.0.1:" + port,
		Check: &consulapi.AgentServiceCheck{
			CheckID: serviceId,
			Status: consulapi.HealthPassing,
			TTL: "1m",
		},
	})
	if err != nil {
		log.Println("can't add service to consul:", err)
		return
	}
	log.Println("register service in consul:", serviceId)

	go func() {
		for {
			log.Println("check ttl")
			err := consul.Agent().PassTTL(serviceId, "check")
			if err != nil {
				log.Println("fail pass ttl")
			}
			<- time.After(30 * time.Second)
		}
	}()

	defer func() {
		err := consul.Agent().ServiceDeregister(serviceId)
		if err != nil {
			log.Println("can't deregister service in consul", err)
			return
		}
		log.Println("deregister service in consul", serviceId)
	}()

	log.Println("starting service at port", port)
	go server.Serve(lis)
	fmt.Scanln()
}
