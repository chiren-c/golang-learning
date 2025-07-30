package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 100
	var ptr *int // 声明指针类型
	ptr = &a     // 初始化指针类型值为变量 a
	fmt.Println("a", a, reflect.TypeOf(a))
	fmt.Println("ptr", ptr, reflect.TypeOf(ptr))
	fmt.Println("*ptr", *ptr, reflect.TypeOf(*ptr))
	// fmt.Printf("%p\n", ptr)
	// fmt.Printf("%d\n", *ptr)
	// v1 := new(int)
	// fmt.Println("v1", *v1, reflect.TypeOf(v1)) // v1 0 *int
	// 通过指针传值
	fmt.Println("------------------")
	b := 200
	swap1(&a, &b)
	fmt.Println(a, b)
	a2 := 10
	b2 := 5
	swap2(a2, b2)
	fmt.Println(a2, b2)
}

func swap1(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func swap2(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}
