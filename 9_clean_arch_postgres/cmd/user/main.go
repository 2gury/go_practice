package main

import (
	"go_practice/9_clean_arch_db/config"
	grpcUser "go_practice/9_clean_arch_db/internal/user/delivery/grpc"
	userRepository "go_practice/9_clean_arch_db/internal/user/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	confg, err := config.LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", confg.GetUserConnectionServerString())
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	postgresConnection, err := confg.Postgres.GetPostgresDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer postgresConnection.Close()

	userRep := userRepository.NewUserPgRepository(postgresConnection)

	server := grpc.NewServer()
	grpcUser.RegisterUserServiceServer(server, grpcUser.NewUserService(userRep))

	log.Printf("starting session service at %s", confg.GetUserConnectionServerString())
	server.Serve(lis)
}
