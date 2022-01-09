package main

import (
	mux2 "github.com/gorilla/mux"
	clean_arch "go_practice/8_clean_arch"
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

	srv := clean_arch.NewServer("8080", mux)
	if err := srv.Run(); err != nil {
		log.Printf("Error while launch server: %s", err)
	}
}
