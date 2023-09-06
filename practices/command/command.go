package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

	//flag包解析

	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "host", "", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	flag.Parse()

	fmt.Printf("user=%v password=%v host=%v port=%v", user, pwd, host, port)
}
