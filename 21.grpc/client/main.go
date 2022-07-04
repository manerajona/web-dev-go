package main

import (
	"context"
	"fmt"
	"github.com/manerajona/web-dev-go/21.grpc/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := echo.NewEchoServiceClient(conn)

	resp, err := cli.Echo(ctx, &echo.EchoRequest{
		Message: "Hello World!",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Got from server:", resp.Response)
}
