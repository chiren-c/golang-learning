package main

import (
	"fmt"
	"reflect"
	"unsafe"
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
	swap1(&a, &b)     // 200 100
	fmt.Println(a, b) // 200 100
	a2 := 10
	b2 := 5
	swap2(a2, b2)       // 5 10
	fmt.Println(a2, b2) // 10 5
	fmt.Println("----------------------")
	// 指针类型转化
	// unsafe.Pointer 是一个万能指针，可以在任何指针类型之间做转化
	// 这就绕过了 Go 的类型安全机制，所以是不安全的操作。
	c1 := 10
	var c2 *int = &c1
	fmt.Println("c2", c2)
	var c3 *float32 = (*float32)(unsafe.Pointer(c2))
	*c3 = *c3 * 10
	fmt.Println("c1", c1, reflect.TypeOf(c1))
	fmt.Println("c3", *c3, reflect.TypeOf(c3))
	fmt.Println("----------------------")
	// 指针运算实现
	// 绕过 Go 指针的安全限制
	// 实现对指针的动态偏移和计算了
	// 这会导致即使发生数组越界了，也不会报错
	// 而是返回下一个内存地址存储的值，这就破坏了内存安全限制
	// 所以这也是不安全的操作
	d1 := [3]int{1, 2, 3}
	d2 := &d1
	// 将指针向前移动一个元素的大小
	d3 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(d2)) + unsafe.Sizeof(d1[0])))
	fmt.Println("uintptr", uintptr(unsafe.Pointer(d2)))
	fmt.Println("Sizeof", unsafe.Sizeof(d1[0]))
	fmt.Println("d3", *d3, reflect.TypeOf(d3))
	*d3 += 10
	fmt.Println("d3", *d3, reflect.TypeOf(d3))
	fmt.Println("d1", d1, reflect.TypeOf(d1))
}

func swap1(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func swap2(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}
