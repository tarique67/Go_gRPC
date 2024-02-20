package main

import (
	"context"
	"fmt"
	hello "hello-world/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := hello.NewGreeterClient(conn)

	resp, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "Tarique"})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
