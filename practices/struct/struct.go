package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//声明 方式1
	var p1 Person
	p1.Name = "Tom"
	p1.Age = 11
	//2 {}里可以自己赋初值
	var p2 Person = Person{}
	p2.Name = "Jeff"
	p2.Age = 12
	//3
	var p3 *Person = new(Person)
	//(*p3).Name="Smith" 等价于 p3.Name="Smith" 底层会给p3加上取值运算
	(*p3).Name = "Smith"
	p3.Age = 10
	//4
	var p4 *Person = &Person{}
	(*p4).Name = "Smith"
	p4.Age = 10

	fmt.Println(p1)

	/*第3、4种返回的是结构体指针
	结构体的所有字段在内存中是连续的
	结构体在和其它类型进行转换的时候，需要有完全相同的字段
	结构体进行type重新定义（相当于取别名），Golang认为是新的数据类型，但是相互间可以强制转换
	struct上每个字段可以写上一个tag，该tag可以通过反射机制进行获取，常见的场景是序列化和反序列化
	*/

	//json.Marshal()使用到了反射机制
	jsonStr, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(string(jsonStr))
	}
}
