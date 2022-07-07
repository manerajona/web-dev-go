package main

import (
	"context"
	"fmt"
	"github.com/manerajona/web-dev-go/18.grpc/echo"
	"google.golang.org/grpc"
	"net"
)

type EchoServer struct{}

func (e *EchoServer) Echo(_ context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Response: "My Echo: " + req.Message,
	}, nil
}

func main() {
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	echo.RegisterEchoServiceServer(s, &EchoServer{})

	fmt.Println("Now serving at port 8080")
	if err = s.Serve(lst); err != nil {
		panic(err)
	}
}
