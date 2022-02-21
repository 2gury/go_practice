package main

import (
	"context"
	"go_practice/17_protobuf/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"time"
)

func timingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error)  {
	start := time.Now()

	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	log.Printf("[METHOD]: %v; [REQ]: %v; [REPLY]: %#v [ERR]: %v; [TIME]: %v; [METADATA]: %v", info.FullMethod, req, reply,  err, time.Since(start), md)

	return reply, err
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(timingInterceptor))
	session.RegisterAuthServer(server, NewSessionManager())

	log.Println("starting server at :8081")
	server.Serve(lis)
}