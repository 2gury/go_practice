package config

import (
	"context"
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
	return s.Server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

