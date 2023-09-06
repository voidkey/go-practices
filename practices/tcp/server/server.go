package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {

	defer conn.Close()

	for {

		//fmt.Printf("Server is waiting for the messages sent by : %v\n ", conn.RemoteAddr().String())
		buf := make([]byte, 1024)
		/*
			conn.Read(buf)
			1.waiting for messages sent by client through conn
			2.if client doesn't write messages, then goroutine will be blocked here
		*/
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("IP:%v disconnect!\n", conn.RemoteAddr().String())
			} else {
				fmt.Println("ERROR = ", err)
			}
			return // Client shutdown will cause this error to occur
		}
		//3.displaying the message sent by the client to the terminal of the server
		fmt.Print(string(buf[:n]))
		fmt.Println("Sent by IP: ", conn.RemoteAddr().String())
	}

	// fmt.Printf("Server received ")
}

func main() {
	fmt.Println("Server begins to listen...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err= ", err)
		return
	}
	defer listen.Close()
	fmt.Println("listen successed = ", listen)

	for {
		fmt.Println("Waiting for the connection of client")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc=%v, client IP=%v\n", conn, conn.RemoteAddr().String())
			go process(conn)
		}
	}

}
