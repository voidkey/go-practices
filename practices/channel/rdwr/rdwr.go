package main

import (
	"fmt"
	"time"
)

func WriteData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("Write:%v\n", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("Read:%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	//time.Sleep(time.Second*10)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
