package main

import (
	"errors"
	"fmt"
)

func test() {
	//defer + recover, recover只有在defer调用的函数中有效, 在未发生panic时调用recover，recover会返回nil
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func redadConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := redadConf("config.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02()继续执行")
}

func main() {
	test()
	test02()
}
