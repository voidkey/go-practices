package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name   string
	gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed")
		return
	}
	//3.渲染模板

	name := "薛卓飞"

	u1 := &User{
		Name:   "Jeffery",
		gender: "male",
		Age:    188,
	}

	m1 := map[string]interface{}{
		"Name":   "Jeffery",
		"gender": "male",
		"Age":    188,
	}

	hobbyList := []string{
		"唱",
		"跳",
		"RAP",
	}

	err = t.Execute(w, map[string]interface{}{
		"name":  name,
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("Execute failed")
		return
	}
}

func praise(w http.ResponseWriter, r *http.Request) {
	//定义一个夸人函数
	praiseFunc := func(name string) (string, error) {
		return name + " is awesome and outstanding!", nil
	}
	//2.解析模板
	t := template.New("praise.tmpl")
	//告诉模板引擎有一个自定义函数，需要放在解析函数之前
	t.Funcs(template.FuncMap{
		"praise": praiseFunc,
	})
	_, err := t.ParseFiles("./praise.tmpl")
	if err != nil {
		fmt.Println("template.New().ParseFiles() failed,err:", err)
		return
	}

	name := "fei"
	//3.渲染模板
	t.Execute(w, name)

}

func nest(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := User{
		Name:   "xzf",
		gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func index(w http.ResponseWriter, r *http.Request) {
	//解析模板
	tmpl, err := template.ParseFiles("./base.tmpl", "./index.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	//渲染模板
	name := "XUE"
	err = tmpl.ExecuteTemplate(w, "index.tmpl", name)
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
	/*
		如果我们的模板名称冲突了，例如不同业务线下都定义了一个index.tmpl模板，我们可以通过下面两种方法来解决。

		1.在模板文件开头使用{{define 模板名}}语句显式的为模板命名。
		2.可以把模板文件存放在templates文件夹下面的不同目录中，然后使用template.ParseGlob("templates/**/ /*.tmpl")解析模板。
	 */
}

func home(w http.ResponseWriter, r *http.Request) {
	//解析模板
	tmpl, err := template.ParseFiles("./base.tmpl", "./home.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	//渲染模板
	err = tmpl.ExecuteTemplate(w, "home.tmpl", "FEI")
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/", praise)
	http.HandleFunc("/nest", nest)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
