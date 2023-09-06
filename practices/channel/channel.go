package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Color string
}

/*
通过全局变量加锁同步来解决goroutine的通讯不完美
	主线程在等待所有goroutine全部完成的时间不好确定，
	通过全局变量加锁同步来实现通讯，也并不利用多个线程对全局变量的读写操作

channel本质是一个队列，FIFO，本身是线程安全的，多goroutine访问时不需要加锁
channel是有类型的，一个string的channel只能存放string类型数据
channel是引用类型，必须初始化（make）后才能写入数据

channel可以声明未只读或只写
	var xxx chan<- int
	var xxx <-chan int

实际开发中，可能不好确定什么时候关闭管道，使用select 可以解决从管道读取数据的阻塞问题
goroutine使用recover，解决协程中出现panic，导致程序崩溃问题
*/
func main() {
	intChan := make(chan int, 3)
	intChan <- 10
	num := 86
	intChan <- num

	//channel加入数据不能超过其容量，否则会死锁。在slice中append可以超过其容量，但底层数组会变换
	fmt.Println(intChan, &intChan)

	//从管道中读取数据,如果管道里没有数据再读取数据，会报错 deadlock
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)

	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "jerry"
	cat := Cat{"Tom", "Grey"}
	allChan <- cat

	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("type=%T,newCat=%v\n", newCat, newCat)
	a := newCat.(Cat) //编译时仍然将newcat看成是一个接口，没办法调用该结构体中的变量,通过类型断言解决
	fmt.Printf("Name=%v\n", a.Name)
}

func channelClose() {
	//channel关闭后不可写入数据，但仍然可以读取数据
	intChan := make(chan int, 3)
	intChan <- 10
	num := 86
	close(intChan)
	intChan <- num
}

func channelTraversal() {
	/*使用for-ranges遍历channel时：
	channel未关闭，遍历会报错 deadlcok
	channel已关闭，遍历正常运行
	*/
	intChan := make(chan int, 3)
	intChan <- 10
	num := 86
	intChan <- num
	close(intChan)

	for v := range intChan {
		fmt.Println(v)
	}

}

func channelRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Wrong")
		}
	}()
	var myMap map[int]string
	myMap[0] = "str"
}
