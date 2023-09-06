package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Character struct {
	Name string
}

//该方法必须Go的RPC规则
//1.方法只能有两个可序列化的参数，其中第二个参数是指针类型
//2.方法要返回一个error类型，同时必须是公开的方法
//如 cahnnel、complex、func均不能进行序列化

func (c *Character) Say(req string, res *string) (err error) {
	*res = "Hello " + req
	fmt.Println(res)
	return
}

func main() {
	//Register
	if err := rpc.RegisterName("xzf", new(Character)); err != nil {
		fmt.Println("Register PRC service failed! ERROR:", err)
		return
	}
	//Listen
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Listen failed! ERROR:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening...")

	//Establish Connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept() failed! ERROR:", err)
			return
		}
		//Bind service
		go rpc.ServeConn(conn)
	}

}
