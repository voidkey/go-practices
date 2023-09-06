package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"greeter/proto/greeter"
	"net"
)

type Hello struct {
}

func (h *Hello) SayHello(ctx context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{
		Message: "Hello! " + req.Name,
	}, nil
}

func main() {
	fmt.Println("")
	//Initial
	grpcServer := grpc.NewServer()
	//Register
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	//Set listener
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("err:", err)
	}
	//Start Server
	grpcServer.Serve(listener)
}
