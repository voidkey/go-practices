package main

import (
	"fmt"
	"net/rpc"
)

type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsRes struct {
	Result  bool
	Message string
}

type GetGoodsReq struct {
	Id int
}

type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func CallHelloDemo() {
	//Connection
	conn, err := rpc.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("rpc.Dial failed! ERROR:", err)
		return
	}
	defer conn.Close()

	//Call Remote Method
	var reply string
	if err = conn.Call("xzf.Say", "Hutao", &reply); err != nil {
		fmt.Println("conn.Call", err)
		return
	}
	fmt.Println(reply)
}

func CallGoodsDemo() {
	//Connection
	conn, err := rpc.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("rpc.Dial failed! ERROR:", err)
		return
	}
	defer conn.Close()

	//Call Remote Method
	request := AddGoodsReq{
		Id:      86,
		Title:   "Staff of Homa",
		Price:   648,
		Content: "又歪了",
	}
	var reply AddGoodsRes
	if err = conn.Call("goods.AddGoods", request, &reply); err != nil {
		fmt.Println("conn.Call", err)
		return
	}
	fmt.Println(reply.Message)
	
	var getRep GetGoodsRes
	if err = conn.Call("goods.GetGoods", GetGoodsReq{
		Id: 77,
	}, &getRep); err != nil {
		fmt.Println("conn.Call", err)
		return
	}
	fmt.Println(getRep)
}

func main() {
	//CallHelloDemo()
	CallGoodsDemo()
}
