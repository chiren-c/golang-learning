package main

import (
	"fmt"
	"time"
)

// 为函数类型设置别名提高代码可读性
type addFuncType func(int, int) int

func main() {
	a := 2
	b := 8
	c := add(a, b)
	fmt.Println("add", c)
	fmt.Println("-------------------")
	v1 := execTime(add)
	fmt.Println("add", v1(a, b))
	fmt.Println("-------------------")
	v2 := execTime(multiply1)
	fmt.Println("multiply1", v2(a, b))
	v3 := execTime(multiply2)
	fmt.Println("multiply2", v3(1, 4))
}
func add(a, b int) int {
	return a + b
}

func multiply1(a, b int) int {
	return a * b
}

func multiply2(a, b int) int {
	return a << b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
func execTime(f addFuncType) addFuncType {
	return func(a, b int) int {
		start := time.Now()      // 起始时间
		c := f(a, b)             // 执行乘法运算函数
		end := time.Since(start) // 函数执行完毕耗时
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return c // 返回计算结果
	}
}
