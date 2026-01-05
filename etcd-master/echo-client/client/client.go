package client

import (
	"context"
	"etcd/echo"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func CallEcho(conn *grpc.ClientConn) {
	client := echo.NewEchoClient(conn)
	in := &echo.EchoRequest{
		Message: "hello server",
	}
	res, err := client.UnaryEcho(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client recv: ", res.Message)
}

func CallEchoLock(conn *grpc.ClientConn) {
	client := echo.NewEchoClient(conn)
	in := &echo.EchoRequest{
		Message: "hello server",
	}
	res, err := client.UnaryLock(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client recv: ", res.Message)
}
