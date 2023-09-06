package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
	进程是程序在os中的一次执行过程，是系统进行 资源分配和调度 的基本单位
	线程是进程的一个执行实例，是 程序执行 的最小单位，是比进程更小的能独立运行的基本单位
	一个进程能创建和销毁多个线程，同一个进程中的多个线程可以并发执行
	一个程序至少有一个进程，一个进程至少有一个线程

	多线程程序在单核上运行，就是并发（在一个cpu上进行轮询操作），程序员宏观上看到像是同时运行，实际上围观同一时刻只有一个线程执行）
	多线程程序在多核上运行，就是并行

	Go协程的特点：
	1.有独立的栈空间
	2.共享程序堆空间
	3.调度由用户控制
	4.协程是轻量级的线程

	主线程是一个物理线程，直接作用在CPU上，是重量级的，非常消耗CPU资源
	协程是从主线程开启的，轻量级的，是逻辑态，资源消耗相对较少

	Golang的协程机制是重要特点，可以轻松开启上万个协程，其它编程语言地带并发机制一般基于线程的，开启过多线程资源消耗也很大


	MPG模式
	M（Machine）：操作系统的主线程（物理线程）
	P（Processor）：协程执行需要的上下文
	G（Goroutine）：协程
*/
func test() {
	for i := 0; i < 100; i++ {
		fmt.Println("hello,test" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	//开启协程，主线程结束，协程也跟着结束
	go test()
	for i := 0; i < 5; i++ {
		fmt.Println("hello,main" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	fmt.Println("CPU数目：", runtime.NumCPU())
	//go1.8之后默认在多核下运行程序，18前需要设置CPU数目
	//runtime.GOMAXPROCS(1)
}
