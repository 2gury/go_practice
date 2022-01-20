package config

import (
	"context"
	"log"
	"net/http"
	"time"
)

type HTTPServer struct {
	Server *http.Server
}

func NewServer(port string, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		Server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  time.Second,
			WriteTimeout: time.Second,
		},
	}
}

func (s *HTTPServer) Run() error {
	log.Printf("starting server at port%s", s.Server.Addr)
	return s.Server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
