package main

import (
    "context"
    "testing"
    pb "hello-world/proto"
)
func TestSayHelloWithName(t *testing.T) {
    s := &Server{}
    ctx := context.Background()
    req := &pb.HelloRequest{Name: "TestName"}
    resp, err := s.SayHello(ctx, req)
    if err != nil {
        t.Errorf("SayHello() error = %v", err)
        return
    }
    expected := "Hello, TestName!"
    if resp.GetMessage() != expected {
        t.Errorf("SayHello() = %v, want %v", resp.GetMessage(), expected)
    }
}
func TestSayHelloWithoutName(t *testing.T) {
    s := &Server{}
    ctx := context.Background()
    req := &pb.HelloRequest{Name: ""}
    resp, err := s.SayHello(ctx, req)
    if err != nil {
        t.Errorf("SayHello() error = %v", err)
        return
    }
    expected := "Hello, !"
    if resp.GetMessage() != expected {
        t.Errorf("SayHello() = %v, want %v", resp.GetMessage(), expected)
    }
}