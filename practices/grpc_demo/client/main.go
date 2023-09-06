package main

import (
	"client/proto/greeter"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	grpcClient, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	client := greeter.NewGreeterClient(grpcClient)
	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "Hutao",
	})
	fmt.Printf("%#v\n", res)
	fmt.Printf("%v\n", res.Message)
}
