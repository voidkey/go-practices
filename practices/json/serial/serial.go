package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "Shrek",
		Age:      100,
		Birthday: "2000-01-01",
		Salary:   10000.0,
		Skill:    "Eat",
	}

	//将struct序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("Serialization failed: ", err)
	}
	fmt.Printf("Result=%v\n", string(data))
}

func testMap() {
	//key是string，value是任意类型
	a := make(map[string]interface{})
	a["name"] = "Godzilla"
	a["age"] = 1000
	a["skill"] = "Dive"

	//map本身是引用类型，不需要指针
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Serialization failed: ", err)
	}
	fmt.Printf("Result=%v\n", string(data))
}

func testSlice() {
	var mapSlice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "Godzilla"
	m1["age"] = 1000
	m1["skill"] = "Dive"
	mapSlice = append(mapSlice, m1)
	m2 := make(map[string]interface{})
	m2["name"] = "Bowser"
	m2["age"] = 99
	m2["skill"] = [2]string{"fire breathing", "breaking"}
	mapSlice = append(mapSlice, m2)

	data, err := json.Marshal(mapSlice)
	if err != nil {
		fmt.Println("Serialization failed: ", err)
	}
	fmt.Printf("Result=%v\n", string(data))
}

func testFloat64() {
	var num float64 = 100.86
	data, err := json.Marshal(num)
	if err != nil {
		fmt.Println("Serialization failed: ", err)
	}
	fmt.Printf("Result=%v\n", string(data))
}

func main() {
	testMap()
	testSlice()
	testFloat64()
}
