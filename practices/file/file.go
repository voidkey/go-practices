package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func openfile() {
	//open file
	file, err := os.Open("D:/GOPATH/go/src/go_code/practices/file/test.txt")
	if err != nil {
		fmt.Println("open file error =", err)
	}
	fmt.Println("file =", file)

	err = file.Close()
	if err != nil {
		fmt.Println("open file error =", err)
	}
}

func readfile() {
	//read file
	file, err := os.Open("D:/GOPATH/go/src/go_code/practices/file/test.txt")
	if err != nil {
		fmt.Println("open file error =", err)
	}
	defer file.Close()
	/*
		方式一：bufio
		创建一个 *Reader，是带缓冲的
		const(defaultBufSize=4096) 默认是4096字节
	*/
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')

		fmt.Println(str)
		if err == io.EOF {
			fmt.Println("EOF =", err)
			break
		}
	}
	/*
		方式二：ioutil.ReadFile
		一次性读取，适用于文件不大的情况
		不用显式的open close文件，因为文件的open和close被封装到ReadFile函数内部
	*/
	content, err := ioutil.ReadFile("D:/GOPATH/go/src/go_code/practices/file/test.txt")
	if err != nil {
		fmt.Println("read error =", err)
	}
	fmt.Printf("%v", string(content)) //[]byte
}

func writefile1() {
	//write file method 1 创建一个新文件并写入数据
	filePath := "D:/GOPATH/go/src/go_code/practices/file/write_test.txt"
	wfile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	str := "hello,fei!\n"
	writer := bufio.NewWriter(wfile)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//调用flush真正将缓冲 []buf里的数据写进文件中
	writer.Flush()
}

func writefile2() {
	//打开一个文件，并写入数据覆盖原内容
	filePath := "D:/GOPATH/go/src/go_code/practices/file/write_test.txt"
	wfile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	str := "こにちは，せつ!\n"
	writer := bufio.NewWriter(wfile)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func writefile3() {
	//打开一个文件，并写入数据追加至原内容
	filePath := "D:/GOPATH/go/src/go_code/practices/file/write_test.txt"
	wfile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	str := "FEI\n"
	writer := bufio.NewWriter(wfile)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func writefile4() {
	//打开一个文件，并写入数据追加至原内容
	filePath := "D:/GOPATH/go/src/go_code/practices/file/write_test.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Printf("%v", str)
		if err == io.EOF {
			fmt.Println("EOF =", err)
			break
		}
	}

	str := "飛\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func read2write() {
	file1path := "D:/GOPATH/go/src/go_code/practices/file/read2write1.txt"
	file2path := "D:/GOPATH/go/src/go_code/practices/file/read2write2.txt"
	data, err := ioutil.ReadFile(file1path)
	if err != nil {
		fmt.Println("read file error=", err)
		return
	}
	err = ioutil.WriteFile(file2path, data, 0666)
	if err != nil {
		fmt.Println("write file error=", err)
		return
	}

}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func copyFile(dst string, src string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(dstFile)

	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {
	// openfile()
	// readfile()
	// writefile1()
	// writefile2()
	// writefile3()
	// writefile4()
	//read2write()
	//res,err := pathExists("D:/a.txt")
	/*
		srcFile := "D:/GOPATH/go/src/go_code/practices/file/hutao_hw.png"
		dstFile := "D:/GOPATH/go/src/go_code/practices/file/hutao_copy.png"
		_, err := copyFile(dstFile, srcFile)
		if err == nil {
			fmt.Println("Copy Successed!")
		} else {
			fmt.Println("Copy Failed! ERROR=", err)
		}
	*/

}
