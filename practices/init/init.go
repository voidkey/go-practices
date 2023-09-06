package main

import (
	"fmt"
	"go_code/practices/init/utils"
)

/*
0.通常在init函数完成初始化工作
1.执行顺序：全局变量->init函数->main函数

*/
var age = test()

func test() int {
	fmt.Println("variable successed!")
	return 1
}

func init() {
	fmt.Println("init successed!")
}

func main() {
	fmt.Printf("Name= %s Age= %d", utils.Name, utils.Age)
}
