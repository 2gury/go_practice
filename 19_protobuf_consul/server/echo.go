package main

import (
	"go_practice/19_protobuf_consul/echo"
	"golang.org/x/net/context"
)

type Echo struct {
	echo.UnimplementedEchoServer
	host string
}

func NewEcho(port string) *Echo {
	return &Echo{
		host: port,
	}
}

func (e *Echo) Say(ctx context.Context, input *echo.Input) (*echo.Output, error) {
	
	return &echo.Output{
		Message: input.Message + " " + e.host,
	}, nil
}

