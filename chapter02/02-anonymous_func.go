package main

import "fmt"

func main() {
	// 1、将匿名函数赋值给变量
	add := func(a, b int) int {
		return a + b
	}
	// 调用匿名函数 add
	fmt.Println("add", add(12, 2))
	fmt.Println("-------------------")
	// 定义时直接调用匿名函数
	func(a, b int) {
		fmt.Println("a + b", a+b)
	}(1, 2)
	fmt.Println("-------------------")
	v1 := 1
	v2 := 2
	// 此时返回的是匿名函数
	addFunc := deferAdd(v1, v2)
	// 这里才会真正执行加法操作
	// 这更好地控制何时执行，延迟执行
	fmt.Println("deferAdd", addFunc())
	fmt.Println("deferAdd1", deferAdd(3, 2)())
	fmt.Println("-------------------")
	var j int = 1
	f := func() {
		var i int = 1
		fmt.Printf("i, j: %d, %d\n", i, j)
		// 引用外部变量 j,会改变外面的值
		j += 1
	}
	f()
	fmt.Println("j after f:", j)
	j += 2
	f()
	fmt.Println("-------------------")
	// 将函数类型作为参数
	// 需要严格指定每个参数和返回值的类型
	// add 函数对应的函数类型是 func(int, int) int。
	func(call func(int, int) int) {
		fmt.Println("call", call(1, 2))
	}(add)
	handleAdd(3, 4, "handleAdd", add)
	handleAdd(3, 4, "subtract", subtract)
}

func deferAdd(a, b int) func() int {
	return func() int {
		return a + b
	}
}

// 将匿名函数作为参数
func handleAdd(a, b int, s string, call func(int, int) int) {
	fmt.Println(s, call(a, b))
}

func subtract(a, b int) int {
	return b - a
}
