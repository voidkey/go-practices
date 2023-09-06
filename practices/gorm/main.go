//package main
//
//import (
//	"fmt"
//)
//
////type Product struct {
////	gorm.Model
////	Code  string
////	Price uint
////}
////
////func main() {
////	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
////	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
////	if err != nil {
////		panic("failed to connect database")
////	}
////
////	// 迁移 schema
////	db.AutoMigrate(&Product{})
////
////	// Create
////	db.Create(&Product{Code: "D42", Price: 100})
////
////	// Read
////	var product Product
////	db.First(&product, 1)                 // 根据整型主键查找
////	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
////
////	// Update - 将 product 的 price 更新为 200
////	db.Model(&product).Update("Price", 200)
////	// Update - 更新多个字段
////	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
////	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
////
////	// Delete - 删除 product
////	db.Delete(&product, 1)
////}
//
//func main() {
//	s := []int{1, 2, 3}
//	fmt.Println(s[1:])
//	fmt.Println(s[1:len(s)])
//	s = s[1:len(s)]
//	fmt.Println(s)
//	s = s[1:len(s)]
//	fmt.Println(s)
//	s = s + s
//	fmt.Println(s)
//}
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
