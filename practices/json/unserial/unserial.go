package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func unMarshalStruct() {
	str := "{\"Age\":99,\"Name\":\"Bowser\",\"Skill\":\"fire breathing\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("Unserialization failed: ", err)
	}
	fmt.Println(monster)
}

func unMarshalMap() {
	//str := "{\"age\":1000,\"name\":\"Godzilla\",\"skill\":\"Dive\"}"
	str := `{"age":1000,"name":"Godzilla","skill":"Dive"}`
	var a map[string]interface{}

	/*
		这里 声明的map不用make，因为make被封装到unmarshal函数里,且参数必须为指针类型的map变量
		（因为只声明了map,未分配内存,unmarshal函数里新分配内存的map变量无法传递给函数外部的map变量）

		Unmarshal may allocates the variable(map, slice, etc.).
		If we pass a map instead of pointer to a map, then the newly allocated map won't be visible to the caller.
	*/
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("Unserialization failed: ", err)
	}
	fmt.Println(a)
}

func unMarshalSlice() {
	str := "[{\"age\":1000,\"name\":\"Godzilla\",\"skill\":\"Dive\"},{\"age\":99,\"name\":\"Bowser\",\"skill\":[\"fire breathing\",\"breaking\"]}]"
	var mapSlice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &mapSlice)
	if err != nil {
		fmt.Println("Unserialization failed: ", err)
	}
	fmt.Println(mapSlice)
}

func main() {
	unMarshalMap()
	unMarshalSlice()
}
