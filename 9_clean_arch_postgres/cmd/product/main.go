package main

import (
	"go_practice/9_clean_arch_db/config"
	grpcProduct "go_practice/9_clean_arch_db/internal/product/delivery/grpc"
	productRepository "go_practice/9_clean_arch_db/internal/product/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	confg, err := config.LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", confg.GetProductConnectionServerString())
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	postgresConnection, err := confg.Postgres.GetPostgresDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer postgresConnection.Close()

	productRep := productRepository.NewProductPgRepository(postgresConnection)

	server := grpc.NewServer()
	grpcProduct.RegisterProductServiceServer(server, grpcProduct.NewProductService(productRep))

	log.Printf("starting session service at %s", confg.GetProductConnectionServerString())
	server.Serve(lis)
}
