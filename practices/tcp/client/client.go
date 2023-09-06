package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.174.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	defer conn.Close()

	fmt.Println("connect successed = ", conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n') //Readbytes
		if err != nil {
			fmt.Println("readString err=", err)
		}

		line1 := strings.Trim(line, " \r\n")
		if line1 == "exit" {
			fmt.Println("Client Over")
			break
		}
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		fmt.Printf("client sent %d bytes data\n", n)
	}

}
