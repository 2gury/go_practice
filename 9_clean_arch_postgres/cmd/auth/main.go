package main

import (
	"go_practice/9_clean_arch_db/config"
	grpcSess "go_practice/9_clean_arch_db/internal/session/delivery/grpc"
	sessRepository "go_practice/9_clean_arch_db/internal/session/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	confg, err := config.LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", confg.GetAuthConnectionServerString())
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	redisConnection, err := confg.Redis.GetRedisDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer redisConnection.Close()

	sessRep := sessRepository.NewSessionRdRepository(redisConnection)

	server := grpc.NewServer()
	grpcSess.RegisterSessionServiceServer(server, grpcSess.NewSessionService(sessRep))

	log.Printf("starting session service at %s", confg.GetAuthConnectionServerString())
	server.Serve(lis)
}
