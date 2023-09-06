package main

import (
	"fmt"
)

/*
defer会将语句及相关的值同时入栈（独立的defer栈）
defer最主要的价值在于，当函数执行完毕后，可以及时的释放函数创建的资源
在函数中,程序员经常需要创建资源(比如:数据库连接、文件句柄、锁等) ,
为了在函数执行完 毕后,及时的释放资源,Go 的设计者提供 defer (延时机制)。
*/
func sum(n1 int, n2 int) int {
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)
	res := n1 + n2
	fmt.Println("n1+n2=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}
