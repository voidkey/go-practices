package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 string = "12345"
	var a int64
	a, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("type = %T, value = %d\n", a, a)

	var str2 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str2)
	fmt.Printf("type = %T, value = %v\n", b, b)

	var str3 string = "3.1415"
	var c float64
	c, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("type = %T, value = %v\n", c, c)

	var str4 string = "12345"
	var d int
	d, _ = strconv.Atoi(str4)
	fmt.Printf("type = %T, value = %d\n", d, d)

	var str5 string = "abc"
	var e []byte = make([]byte, 3)
	e = []byte(str5)
	fmt.Printf("type = %T, value = %d\n", e, e)

}
