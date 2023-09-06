package main

import (
	"fmt"
	"testing"
)

/*
testing框架
1.将 xxx_test.go文件引入import
2.在 main函数里调用所有TestXxx函数
*/

func TestAddUpper(t *testing.T) {
	res := AddUpper1(10)
	if res != 55 {
		t.Fatalf("ERROR,result=%v", res)
	}
	t.Logf("Right!")
}

func TestHi(t *testing.T) {
	fmt.Println("Hi!")
}

/*
1.go test 运行错误才有日志 go test -v 都有日志
2.PASS 表示 测试成功 FAIL 表示 失败
3.测试单个文件 go test xxx_test.go xxx.go
4.测试单个方法 go test -test.run TestXxx
*/
