package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]float64, 10)

	lock sync.Mutex // lock全局的互斥锁 Mutex互斥
)

func test(n int) {
	var res float64 = 1
	for i := 1; i <= n; i++ {
		res *= float64(i)
	}
	lock.Lock()
	myMap[n] = res //error: concurrent map writes
	lock.Unlock()
}

/*
在编译一个程序时加入 -race参数可以知道是否存在资源竞争问题
不同goroutine之间如何通信：
1.使用全局变量加锁同步,等待的协程会进入队列
2.channel
*/
func main1() {
	//开启协程完成

	for i := 1; i <= 10; i++ {
		go test(i)
	}
	time.Sleep(time.Second)
	lock.Lock()
	for i := 1; i <= 10; i++ {
		fmt.Println(myMap[i])
	}
	lock.Unlock()

}
