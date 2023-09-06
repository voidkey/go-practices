package main

import "fmt"

func main() {
	res := sum(10, 1, 2, 3, 4, 5)
	fmt.Println("res=", res)

}

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
