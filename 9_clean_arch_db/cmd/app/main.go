package main

import (
	"fmt"
	mux2 "github.com/gorilla/mux"
	"go_practice/8_clean_arch/config"
	"go_practice/8_clean_arch/internal/product/delivery"
	"go_practice/8_clean_arch/internal/product/repository"
	"go_practice/8_clean_arch/internal/product/usecases"
	"log"
)

func main() {
	confg, _ := config.LoadConfig("./config.json")
	fmt.Println(confg)
	productRepository := repository.NewProductArrayRepository()
	productUsecase := usecases.NewProductUsecase(productRepository)
	productHandler := delivery.NewProductHandler(productUsecase)

	mux := mux2.NewRouter()
	productHandler.Configure(mux)

	srv := config.NewServer("8080", mux)
	log.Fatal(srv.Run())
}
