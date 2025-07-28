package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a1 string = "Hello, World!"
	a2 := "世界和平"
	a3 := `反引号来定义一行'字'符串 "Golang":`
	a4 := `多行字符串
	可以使用反引号来定义
	多行字符串，反引号内的内容会原样保留，包括换行符和空格。
`
	a5 := "字符串拼接" + "示例"
	a6 := a1[0]   // 获取索引 0 的子串
	a7 := a1[:5]  // 获取索引 5（不含）之前的子串
	a8 := a1[3:6] // 获取索引 3（不含）到 6（含）的子串
	a9 := a1[6:]  // 获取索引 6 （含）之后的子串

	fmt.Printf("a1: %s\n", a1)
	fmt.Printf("a2: %s\n", a2)
	fmt.Printf("a3: %s\n", a3)
	fmt.Printf("a4: %s\n", a4)
	fmt.Printf("a5: %s\n", a5)
	fmt.Printf("a6: %c\n", a6) // 输出字符 H
	fmt.Printf("a7: %s\n", a7)
	fmt.Printf("a8: %s\n", a8)
	fmt.Printf("a9: %s\n", a9)
	fmt.Println("----------------------")
	msg := "Hello, 世界"
	l := len(msg) // 获取到的是 msg 的字节长度
	// byte 代表 UTF-8 编码中单个字节的值（它也是 uint8 类型的别名，两者是等价的，因为正好占据 1 个字节的内存空间,中文字符在 UTF-8 中占 3 个字节）
	for i := 0; i < l; i++ {
		ch := msg[i] // 依据下标取字符串中的字符，ch 类型为 byte
		// UTF-8 编码不能 string 这样转化，英文字符没问题，因为一个英文字符就是一个字节，中文字符则会乱码，因为一个中文字符编码需要三个字节，转化单个字节会出现乱码。
		fmt.Println(i, ch, string(ch), reflect.TypeOf(ch))
	}
	fmt.Println("----------------------")
	//  Unicode 字符遍历。rune 代表单个 Unicode 字符（它也是 int32 类型的别名，因为正好占据 4 个字节的内存空间）
	for i, ch := range msg {
		fmt.Println(i, ch, string(ch), reflect.TypeOf(ch))
	}
}
