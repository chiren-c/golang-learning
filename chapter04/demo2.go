package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	lock := sync.WaitGroup{}
	// lock.Add(10) // 设置等待组计数为10
	for i := 0; i < 10; i++ {
		lock.Add(1) // 每次添加一个计数
		go add3(1, i, lock.Done)
	}
	lock.Wait()
	end := time.Now()
	fmt.Printf("程序执行耗时(s)\n", end.Sub(start).Seconds())
}

func add3(a, b int, f func()) {
	defer func() {
		f()
	}()
	c := a + b
	fmt.Printf("add3: %d + %d = %d\n", a, b, c)
}
