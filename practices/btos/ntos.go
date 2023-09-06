package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 1
	var b float32 = 3.1415926
	var c bool = true
	var d byte = 'd'
	var e []byte = []byte{'e', 'e', 'e'}
	var f byte = 'f'
	var g int = 99
	var h int64 = 123
	var i float64 = 2.712828
	var str string

	str = fmt.Sprintf("%d", a)
	fmt.Printf("type = %T, value = %q\n", str, str)

	str = fmt.Sprintf("%f", b)
	fmt.Printf("type = %T, value = %v\n", str, str)

	str = fmt.Sprintf("%t", c)
	fmt.Printf("type = %T, value = %s\n", str, str)

	str = fmt.Sprintf("%c", d)
	fmt.Printf("type = %T, value = %s\n", str, str)

	str = string(e)
	fmt.Printf("type = %T, value = %s\n", str, str)

	str = string(f)
	fmt.Printf("type = %T, value = %s\n", str, str)

	str = strconv.FormatInt(int64(g), 10)
	fmt.Printf("type = %T, value = %q\n", str, str)

	str = strconv.Itoa(int(h))
	fmt.Printf("type = %T, value = %s\n", str, str)

	str = strconv.FormatFloat(i, 'f', 10, 64)
	fmt.Printf("type = %T, value = %q\n", str, str)
}
