package main

import "fmt"

func main() {

	var slice []int = make([]int, 10)

	for i := 0; i < len(slice); i++ {
		slice[i] = 5 * i
	}
	fmt.Println(slice)
	printSlice(slice)

}
func printSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}
	fmt.Printf("\nThe length of slice is %d\n", len(slice))
	fmt.Printf("The capacity of slice is %d\n", cap(slice))
}
