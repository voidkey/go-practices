package main

import (
	binary "go_code/practices/sort_retrieve/binary_retrieve"
	bubble "go_code/practices/sort_retrieve/bubble_sort"
)

func main() {
	data1 := make([]int, 10)
	data1 = []int{1, 3, 5, 7, 6, 4, 8, 2, 3, 0}

	bubble.Bubble_sort(data1)
	binary.Binary_Retrieve(data1, 0, 9, 9)
}
