package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func main() {
	// for i := 1; i < 10; i++ {
	// 	go add(1, i)
	// }
	// time.Sleep(1 * time.Second)
	start := time.Now()
	// 避免 counter 值被污染,引入了锁机制
	lock := &sync.Mutex{}
	for i := 1; i < 10; i++ {
		go add2(1, i, lock)
	}
	// 主协程中通过一个死循环来判断 counter 的值，
	// 只有当它大于等于 10 时，才退出循环，进而退出整个程序
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched() // 让出当前 goroutine 的执行权
		if c >= 9 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

func add(a, b int) int {
	var c = a + b
	fmt.Printf("add(%d, %d) = %d\n", a, b, c)
	go multiply(a, b)
	return c
}

func add2(a, b int, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	var c = a + b
	counter++
	fmt.Printf("add2(%d, %d) = %d\n", a, b, c)
}

func multiply(a, b int) int {
	var c = a * b
	fmt.Printf("multiply(%d, %d) = %d\n", a, b, c)
	return c
}
