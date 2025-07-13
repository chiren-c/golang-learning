package main

import (
	"fmt"
)

func main() {
	// ------------------------------------
	// 预定义常量义示例
	// ------------------------------------
	const (
		// iota 是一个预定义的常量计数器
		e1 = iota     // e1 = 0
		e2 = iota * 2 // e2 = 2
		e3 = iota     // e3 = 2
	)
	const e4 = iota // e4 = 0, iota 在每个 const 块中重置为 0
	const e5 = iota // e5 = 0, iota 在每个 const 块中重置为 0
	// 枚举表示法
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Printf("e1: %v\n", e1)
	fmt.Printf("e2: %v\n", e2)
	fmt.Printf("e3: %v\n", e3)
	fmt.Printf("e4: %v\n", e4)
	fmt.Printf("e5: %v\n", e5)
	fmt.Printf("Sunday: %v\n", Sunday)
	fmt.Printf("Monday: %v\n", Monday)
	fmt.Printf("Tuesday: %v\n", Tuesday)
	fmt.Printf("Wednesday: %v\n", Wednesday)
	fmt.Printf("Thursday: %v\n", Thursday)
	fmt.Printf("Friday: %v\n", Friday)
	fmt.Printf("Saturday: %v\n", Saturday)
	fmt.Printf("numberOfDays: %v\n", numberOfDays)
	fmt.Printf("------------------------------------\n")
}
