package config

import (
	"context"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(port string, mux http.Handler) *HttpServer {
	return &HttpServer{
		server: &http.Server{
			Addr: ":" + port,
			Handler: mux,
			ReadTimeout: time.Second,
			WriteTimeout: time.Second,
		},
	}
}

func (s *HttpServer) Start() error {
	log.Printf("Start server at %s port", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *HttpServer) Shutdown() error {
	return s.server.Shutdown(context.Background())
}
