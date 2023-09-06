package main

func AddUpper1(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

// func main() {
// 	//传统方法 在main函数调用测试
// 	res := AddUpper1(10)

// 	if res == 55 {
// 		fmt.Println("Right")
// 	} else {
// 		fmt.Println("Wrong")
// 	}

// }
