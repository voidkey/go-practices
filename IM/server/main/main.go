package main

import (
	"IM/server/model"
	"fmt"
	"net"
	"time"
)

var (
	listenIp   = "0.0.0.0"
	listenPort = "8889"
)

//Deal with the communication between sever and client
func process(conn net.Conn) {
	defer conn.Close()
	//Read the message sent by client
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Dispatcher()
	if err != nil {
		fmt.Println("Communication error:", err)
		return
	}

}

func init() {
	err := initClient("localhost:6379", 16, time.Second*300)
	if err != nil {
		fmt.Println("initClient err:", err)
	}
	fmt.Println("Init successful!")
	model.MyUserDao = model.NewUserDao(rdb, ctx)
}

func main() {
	fmt.Printf("Server is listening at Port: %v\n", listenPort)
	listenAddr := listenIp + ":" + listenPort
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Listen failed! Error:", err)
		return
	}
	defer listen.Close()
	fmt.Println("Listen successful!")

	fmt.Println("Waiting for the connection of Client...")
	//Waiting for the connection of Client
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() failed! Error:", err)
		}
		go process(conn)
	}

}
