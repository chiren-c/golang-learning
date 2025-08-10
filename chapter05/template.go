package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmp, err := template.ParseFiles("chapter05/hello.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	user := UserInfo{
		Name:   "shu1",
		Gender: "男",
		Age:    18,
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmp.Execute(w, user)
}
