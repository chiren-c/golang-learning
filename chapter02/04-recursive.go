package main

import (
	"fmt"
	"reflect"
	"time"
)

type fibonacciFunc func(int) int

const MAX = 50

var fibs [MAX]int

func main() {
	fmt.Println(fibs, reflect.TypeOf(fibs))
	n := 5
	v1 := execTime2(fibonacci2)
	num := v1(n)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n, num)
}

func execTime2(f fibonacciFunc) fibonacciFunc {
	return func(n int) int {
		s := time.Now()
		c := f(n)
		e := time.Since(s)
		fmt.Printf("--- 执行耗时: %v ---\n", e)
		return c

	}
}

func fibonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci2(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	index := n - 1
	if fibs[index] != 0 {
		return fibs[index]
	}
	num := fibonacci2(n-1) + fibonacci2(n-2)
	fibs[index] = num
	return num
}
