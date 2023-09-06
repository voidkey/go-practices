package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

/*
	Golang的方法作用在指定的数据类型上，因此自定义类型都可以有方法，不只是struct，例如int,float32都可以有

	方法的调用和传参机制和函数基本一致，不同的地方在于方法调用时，会将调用方法的变量，当作实参也传递给方法
	（如果该变量是值类型，则是值拷贝，否则为地址拷贝）

	struct是值传递，如果想通过方法修改结构体变量的值，可以通过结构体指针的方式来处理

	如果一个类型实现了自己的String方法，则fmt.Println()默认调用该变量的String方法进行输出
*/
func (p Person) test() {
	p.Name = "xzf"
	fmt.Println(p.Name)
}

func main() {
	p1 := Person{"tom", 1}

	p1.test()
}

/*
方法和函数的区别
	方法调用方式：变量.方法名（）
	函数调用方式： 函数名（）

	函数，接收者为值类型时，不能将指针类型的数据作为参数直接传递，反之亦然
	方法，可以（&p.test(),虽然变量是地址，但实际上仍然是值传递）
*/
