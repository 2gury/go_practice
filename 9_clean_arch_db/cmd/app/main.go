package main

import (
	mux2 "github.com/gorilla/mux"
	"go_practice/8_clean_arch/config"
	"go_practice/8_clean_arch/internal/product/delivery"
	"go_practice/8_clean_arch/internal/product/repository"
	"go_practice/8_clean_arch/internal/product/usecases"
	"log"
)

func main() {
	confg, _ := config.LoadConfig("./config.json")
	dbConnection, err := confg.Database.GetPostgresDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	productRepository := repository.NewProductPgRepository(dbConnection)
	productUsecase := usecases.NewProductUsecase(productRepository)
	productHandler := delivery.NewProductHandler(productUsecase)

	mux := mux2.NewRouter()
	productHandler.Configure(mux)

	srv := config.NewServer("8080", mux)
	log.Fatal(srv.Run())
}
