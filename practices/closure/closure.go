package main

import (
	"fmt"
	suffix "go_code/practices/closure/makesuffix"
)

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	//闭包是一个函数和与其相关的引用环境组合成的一个整体（实体）

	/*
		AddUpper返回的是一个匿名函数，这个匿名函数引用到函数外的n，
		因此这个匿名函数和n组成一个整体，构成闭包
		闭包是类，函数是操作，n是字段
		当反复调用f函数时，因为n是初始化一次，所以每调用一次就进行累加
	*/
	//累加器
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	suffix.SuffixTest()
}
