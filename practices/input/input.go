package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Printf("name = %v  age = %v\n", name, age)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("name = %v  age = %v", name, age)
}
