package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct {
	Name string
}

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

//该方法必须Go的RPC规则
//1.方法只能有两个可序列化的参数，其中第二个参数是指针类型
//2.方法要返回一个error类型，同时必须是公开的方法
//如 cahnnel、complex、func均不能进行序列化

func (g *Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	fmt.Println(req)
	*res = AddGoodsRes{
		Result:  true,
		Message: "Item Added!",
	}
	return nil
}

func (g *Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	fmt.Println(req)
	*res = GetGoodsRes{
		Id:      77,
		Title:   "AMOS",
		Price:   6480,
		Content: "定轨拉满",
	}
	return nil
}

func main() {
	//Register
	if err := rpc.RegisterName("goods", new(Goods)); err != nil {
		fmt.Println("Register PRC service failed! ERROR:", err)
		return
	}
	//Listen
	listener, err := net.Listen("tcp", "localhost:9000")
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
