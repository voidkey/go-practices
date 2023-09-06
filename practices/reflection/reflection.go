package main

import (
	"fmt"
	"reflect"
)

/*
	反射可以在运行时动态获取变量的各种信息，比如类型type，类别kind
	如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
	通过反射，可以修改变量的值，可以调用关联的方法
	使用反射，需要import("reflect")

	1.relect.Value.kind 获取变量的类别，返回的是一个常量
	2.Type 和 Kind 可能相同也可能不同 比如 var stu Student stu的type是 包名.Student kind是struct
	3.变量、interface()、reflect.Value是可以相互转换的
	4.relect.value类型是运行时的类型，不能直接使用
	5.通过反射来修改变量，注意当使用SetXxx方法来设置需要通过对应的指针类型来完成，这样才能改变传入的变量的值，同时需要使用reflect.Value.Elem()
*/
type Student struct {
	Name string
	Age  int
}

func testInt(b interface{}) {
	//变量、interface()、reflect.Value是可以相互转换的，在使用反射的过程中，通常的方式是变量->接口->reflect reflect->接口->变量
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v, Type of rVal = %T\n", rVal, rVal)
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	num := 10 + rVal.Int()
	fmt.Println("num = ", num)

	iVal := rVal.Interface()
	v := iVal.(int)
	fmt.Println("value=", v)
}

func testStruct(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v, Type of rVal = %T\n", rVal, rVal)
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//反射类型本质是运行时的类型，编译阶段编译器无法确认类型
	fmt.Println("rType=", rType)
	//rKind := rType.Kind()
	rKind := rVal.Kind()
	fmt.Println("rKind=", rKind)

	iVal := rVal.Interface()

	//v,ok := iVal.(Student)
	v := TypeJudge(iVal)

	fmt.Printf("value=%v ,type=%T ", v, v)
}

func testElem(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v\nType = %T\nKind=%v\n", rVal, rVal, rVal.Kind())
	fmt.Printf("bVal = %v\nbType = %T\n", b, b)
	rVal.Elem().SetInt(20)                           //获取指针所指向的变量的值
	fmt.Printf("result=%v\nb= %v\n", rVal.Elem(), b) //不能直接用*rVal,*b,因为在运行时候才知道b是一个指针，编译阶段无法确认
}

func TypeJudge(iVal interface{}) interface{} {
	switch iVal.(type) {
	case Student:
		return iVal.(Student)
	case int:
		return iVal.(int)
	default:
		fmt.Println("wrong type")
	}
	return nil
}

func main() {

	stu := Student{
		Name: "xzf",
		Age:  18,
	}

	testStruct(stu)

	const (
		a = iota //返回该常量在const（）里的行数
		b
		c, d = iota, iota
	)
	fmt.Println(a, b, c, d)

	num1 := 10
	// var num2 *int = new(int)
	// *num2 = 1
	// fmt.Println(num2, *num2)
	testElem(&num1)
	fmt.Printf("numVal = %v\nnumType = %T\n", num1, num1)

	/*
		Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
		如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
	*/
	str := "string"
	rVal := reflect.ValueOf(&str) //这里没有通过接口，需要使用elem就需要传递指针，否则会panic
	fmt.Printf("rVal = %v\nType = %T\nKind=%v\n", rVal, rVal, rVal.Kind())
	rVal.Elem().SetString("hello")
	fmt.Printf("rVal = %v\nType = %T\nKind=%v\n", rVal, rVal, rVal.Kind())
}
