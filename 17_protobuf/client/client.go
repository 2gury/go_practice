package main

import (
	"context"
	"go_practice/17_protobuf/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func timingInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("[METHOD]: %v; [REQ]: %v; [REPLY]: %#v [ERR]: %v; [TIME]: %v; ", method, req, reply,  err, time.Since(start))
	return err
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(timingInterceptor),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	auth := session.NewAuthClient(conn)

	md := metadata.Pairs(
		"key1", "lol",
		"key2", "kek",
	)
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)

	var header, trailer metadata.MD

	sessId, _ := auth.Create(ctx,
		&session.Session{
			Login: "keko",
			UserAgent: "sheesh",

		},
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	log.Printf("[HEADER]: %v; [TRAILER]: %v", header, trailer)

	log.Println(sessId)

	_, err = auth.Check(context.Background(), sessId)
	log.Println(err)

	_, err = auth.Delete(context.Background(), sessId)

	_, err = auth.Check(context.Background(), sessId)
	log.Println(err)
}