package main

import (
	mux2 "github.com/gorilla/mux"
	"go_practice/8_clean_arch/config"
	"go_practice/8_clean_arch/internal/product/delivery"
	repository "go_practice/8_clean_arch/internal/product/repository"
	"go_practice/8_clean_arch/internal/product/usecases"
	"log"
)

func main() {
	//data := repository.NewLocalRepository()
	//
	//rep := repository.NewRepository(data)
	//svc := service.NewService(rep)
	//hnd := handler.NewHandler(svc)

	productRepository := repository.NewProductArrayRepository()
	productUsecase := usecases.NewProductUsecase(productRepository)
	productHandler := delivery.NewProductHandler(productUsecase)

	mux := mux2.NewRouter()
	productHandler.Configure(mux)

	srv := config.NewServer("8080", mux)
	log.Fatal(srv.Run())
}
