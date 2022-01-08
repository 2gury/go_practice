package main

import (
	clean_arch "go_practice/8_clean_arch"
	"go_practice/8_clean_arch/pkg/handler"
	"go_practice/8_clean_arch/pkg/repository"
	"go_practice/8_clean_arch/pkg/service"
	"log"
)

func main() {
	data := repository.NewLocalRepository()

	rep := repository.NewRepository(data)
	svc := service.NewService(rep)
	hnd := handler.NewHandler(svc)

	srv := clean_arch.NewServer("8080", hnd.InitRoutes())
	if err := srv.Run(); err != nil {
		log.Printf("Error while launch server: %s", err)
	}
}
