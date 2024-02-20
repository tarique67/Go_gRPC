package main

import (
	"context"
	"fmt"
	hello "hello-world/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	hello.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &Server{})
	fmt.Println("Port listening on 8081")
	s.Serve(lis)
}
