package main

import (
	"fmt"
)

var (
	Func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//在定义匿名函数时就直接调用，这种方式只能调用一次
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res1)

	//将匿名函数赋给一个变量，再通过该变量调用匿名函数
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(20, 20)
	fmt.Println("res=", res2)

	res3 := Func1(2, 30)
	fmt.Println("res=", res3)

}
