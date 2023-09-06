package bubble_sort

import "fmt"

func Bubble_sort(data []int) {
	temp := 0
	isSorted := true
	fmt.Println("Data: ", data)
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				isSorted = false
				temp = data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
		if isSorted {
			break
		}
	}
	fmt.Println("Result: ", data)
}
