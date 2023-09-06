package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name   string `json:"name"`
	Age    int    `json:"monster_age"`
	Gender string
}

func (m *Monster) Print() {
	fmt.Println(m)
}

func (m *Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func (m *Monster) Set(name string, age int, gender string) {
	m.Name = name
	m.Age = age
	m.Gender = gender
}

func TestStruct(b interface{}) {
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("except struct")
		return
	}
	num := val.NumField()
	fmt.Printf("struct has %v fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值=%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag=%v\n", i, tagVal)
		}
	}

	numMethod := val.NumMethod()
	fmt.Printf("struct has %v methods\n", numMethod)

	val.Method(1).Call(nil) //按照方法的名字字典序排序，所以第二个方法是Print（）

	//call first mehtod
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

}

func main() {
	var a Monster = Monster{
		Name:   "Shrek",
		Age:    18,
		Gender: "male",
	}
	TestStruct(a)
}
