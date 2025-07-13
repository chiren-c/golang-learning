package main

import "fmt"

func main() {
	var a1 string = "Hello, World!"
	a2 := "世界和平"
	a3 := `反引号来定义一行字符串 "Golang":`
	a4 := `多行字符串
	可以使用反引号来定义
	多行字符串，反引号内的内容会原样保留，包括换行符和空格。
`
	a5 := "字符串拼接" + "示例"
	a6 := a1[0]

	fmt.Printf("a1: %s\n", a1)
	fmt.Printf("a2: %s\n", a2)
	fmt.Printf("a3: %s\n", a3)
	fmt.Printf("a4: %s\n", a4)
	fmt.Printf("a5: %s\n", a5)
	fmt.Printf("a6: %c\n", a6) // 输出字符 '示'
}
