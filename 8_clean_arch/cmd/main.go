package main

import (
	clean_arch "go_practice/8_clean_arch"
	"go_practice/8_clean_arch/pkg/handler"
	"log"
)

func main() {
	h := handler.Handler{}

	srv := clean_arch.NewServer("8080", h.InitRoutes())
	if err := srv.Run(); err != nil {
		log.Printf("Error while launch server: %s", err)
	}
}
