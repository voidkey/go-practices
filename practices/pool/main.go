package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func dog(dogChan chan struct{}, catChan chan struct{}) {
	for {
		select {
		case <-dogChan:
			fmt.Println("dog")
			catChan <- struct{}{}
			break
		default:
			break
		}
	}
}

func cat(catChan chan struct{}, fishChan chan struct{}) {
	for {
		select {
		case <-catChan:
			fmt.Println("cat")
			fishChan <- struct{}{}
			break
		default:
			break
		}
	}
}

func fish(fishChan chan struct{}, dogChan chan struct{}) {
	i := 0
	for {
		select {
		case <-fishChan:
			fmt.Println("fish")
			i++ // 计数，打印完之后就溜溜结束了。
			if i > 9 {
				wg.Done()
				return
			}
			dogChan <- struct{}{}
			break
		default:
			break
		}
	}
}

func worker(jobs chan int, res chan int) {
	for v := range jobs {
		res <- v * v
	}
}

func main() {
	jobs := make(chan int, 100)
	res := make(chan int, 100)
	for i := 0; i < 3; i++ {
		go worker(jobs, res)
	}
	for j := 1; j <= 100; j++ {
		jobs <- j
	}

	close(jobs)
	var v int
	for i := 0; i < 100; i++ {
		v = <-res
		fmt.Println(v)
	}

}
