package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Printf("年=%v\n月=%v %v\n日=%v\n", now.Year(), now.Month(), int(now.Month()), now.Day())

	//格式化日期和时间
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-1-2"))
	fmt.Println()

	//时间常量结合休眠

	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}

	//unix unixnano 时间戳，用于获取随机数
	fmt.Println(now.Unix(), now.UnixNano())

	//计时
	/*
		start:=time.Now().Unix()
		function()
		end:=time.Now().Unix()
		time: end-start
	*/

}
