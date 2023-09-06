package bubble_sort

import "fmt"

func Binary_Retrieve(data []int, leftIndex int, rightIndex int, val int) {
	if leftIndex > rightIndex {
		fmt.Println("Failed to find")
		return
	}
	middle := (leftIndex + rightIndex) / 2

	if data[middle] > val {
		Binary_Retrieve(data, leftIndex, middle-1, val)
	} else if data[middle] < val {
		Binary_Retrieve(data, middle+1, rightIndex, val)
	} else {
		fmt.Println("Find data!")
	}

	fmt.Println("Numbers of Retrieve")
}
