package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string = "golang哈" //golang编码 utf-8 汉字占3个字节
	fmt.Println("str len=", len(str))

	//字符串遍历
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("character=%c , %c\n", r[i], str[i])
	}

	//字符串转整数
	a, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("Wrong ", err)
	} else {
		fmt.Println(a)
	}

	//整数转字符串
	b := strconv.Itoa(200)
	fmt.Println(b)

	//字符串转[]byte
	var c = []byte("hello")
	fmt.Printf("%v\n", c)

	//[]byte转字符串
	d := string([]byte{97, 98, 99})
	fmt.Println(d)

	//10进制转2/8/16进制
	e := strconv.FormatInt(123, 2)
	fmt.Println(e)
	e = strconv.FormatInt(123, 8)
	fmt.Println(e)
	e = strconv.FormatInt(123, 16)
	fmt.Println(e)

	//查找子串
	f := strings.Contains("seafood", "foo")
	fmt.Println(f)

	//统计字串数目
	g := strings.Count("cheese", "e")
	fmt.Println(g)

	//不区分大小写的字符串比较
	h := strings.EqualFold("abc", "Abc")
	fmt.Println(h)

	//返回子串第一次出现的索引,不存在返回-1
	i := strings.Index("abefabc", "abc")
	fmt.Println(i)

	//返回子串最后一次出现的索引
	j := strings.LastIndex("abcfabc", "bc")
	fmt.Println(j)

	//替换指定子串,第四个参数为替换数量，-1代表全部替换
	k := strings.Replace("abcabbcab", "ab", "00", 2)
	fmt.Println(k)
	k = strings.Replace("abcabbcab", "ab", "00", -1)
	fmt.Println(k)

	// 按照指定字符，将字符串分割为字符串数组
	l := strings.Split("hello,world,go", ",")
	fmt.Println(l)

	//字符串字母大小写转换
	str = "goLang Hello"
	m := strings.ToLower(str)
	fmt.Println(str, m)
	m = strings.ToUpper(str)
	fmt.Println(str, m)

	//去除字符串两边空格
	n := strings.TrimSpace("[ go lang ]")
	fmt.Println(n)

	//去除字符串两边指定字符
	o := strings.Trim("! go lang! ", " !")
	fmt.Println(o)

	//去除字符串单边指定字符
	p := strings.TrimLeft("! go lang! ", " !")
	fmt.Println(p)
	p = strings.TrimRight(" ! go lang! ", " !")
	fmt.Println(p)

	//判断字符串是否以指定字符串开头/结束
	q := strings.HasPrefix("ftp://192.168.1.1", "ftp")
	fmt.Println(q)
	q = strings.HasSuffix("ftp://192.168.1.1", "10.1")
	fmt.Println(q)

}
