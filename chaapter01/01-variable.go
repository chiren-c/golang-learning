package main

import (
	"fmt"
)

func main() {
	// ------------------------------------
	// 变量声明示例
	// ------------------------------------
	var v1 int // 整型
	var (
		v2 string // 字符串类型
		v3 bool   // 布尔类型
	)
	var v4 [10]int  // 数组类型
	var v5 []int    // 切片类型
	var v6 struct { // 结构体类型
		name string
		age  int
	}
	var v7 *int            // 指针类型
	var v8 map[string]int  // map（字典），key为字符串类型，value为整型
	var v9 func(a int) int // 函数，参数类型为整型，返回值类型为整型

	// 打印上述变量值
	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("v2: %v\n", v2)
	fmt.Printf("v3: %v\n", v3)
	fmt.Printf("v4: %v\n", v4)
	fmt.Printf("v5: %v\n", v5)
	fmt.Printf("v6: %v\n", v6)
	fmt.Printf("v7: %v\n", v7)
	fmt.Printf("v8: %v\n", v8)
	fmt.Printf("v9: %v\n", v9)
	fmt.Printf("------------------------------------\n")
	// ------------------------------------
	// 初始化示例
	// ------------------------------------
	var a1 int
	a1 = 10         // 初始化整型变量
	var a2 int = 20 // 使用类型声明初始化整型变量
	// 注意：简短声明只能在函数内部使用，不能在包级别使用
	a3 := 1                               // 使用简短声明初始化整型变量，:= 运算符左侧的变量应该是未声明过的
	var a4 [5]int = [5]int{1, 2, 3, 4, 5} // 初始化数组
	a5 := [5]int{1, 2, 3}                 // 使用简短声明初始化数组,剩余元素会自动初始化为零值
	// 注意：切片的长度不需要在声明时指定
	// 切片是动态数组，可以根据需要扩展长度
	// 切片的零值是 nil
	a6 := []int{1, 2, 3, 4, 5}           // 使用简短声明初始化切片
	a7 := make([]int, 5)                 // 使用 make 函数初始化切片，长度为5，元素为零值
	a8 := make(map[string]int)           // 使用 make 函数初始化 map，key 为字符串类型，value 为整型
	a9 := map[string]int{"a": 1, "b": 2} // 使用简短声明初始化 map，key 为字符串类型，value 为整型
	var a10 *int                         // 指针类型变量声明
	a10 = new(int)                       // 使用 new 函数分配内存，v7 指向一个整型变量的地址
	*a10 = 42                            // 给指针指向的整型变量赋值

	// 结构体的初始化
	a11 := struct {
		name string
		age  int
	}{
		name: "chen",
		age:  30,
	}
	// 使用 new 函数来初始化结构体
	// new 函数返回一个指向新分配的结构体的指针
	// 结构体的字段可以通过指针访问
	// 注意：new 函数返回的是指针类型
	// 结构体的字段可以通过指针访问
	// 例如：(*a12).name = "chen2"
	a12 := new(struct {
		name string
		age  int
	})
	(*a12).name = "chen2"
	(*a12).age = 31

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a5: %v\n", a5)
	fmt.Printf("a6: %v\n", a6)
	fmt.Printf("a7: %v\n", a7)
	fmt.Printf("a8: %v\n", a8)
	fmt.Printf("a9: %v\n", a9)
	fmt.Printf("a10: %v\n", a10)
	fmt.Printf("a11: %v\n", a11)
	fmt.Printf("a12: %v\n", a12)
	fmt.Printf("------------------------------------\n")

	// ------------------------------------
	// 多重赋值示例
	// ------------------------------------
	var b1, b2 int // 声明两个整型变量
	b1 = 10        // 整型变量赋值
	b2 = 20        // 整型变量赋值
	// 多重赋值
	b1, b2 = b2, b1 // 交换两个整型变量的值(b1 变为 20, b2 变为 10)
	// 注意：多重赋值可以在函数内部使用，也可以在包级别使用
	// 注意：多重赋值可以用于不同类型的变量，但必须保证
	//       左侧变量的类型与右侧变量的类型相同
	//       否则会报编译错误
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("------------------------------------\n")

	// ------------------------------------
	// 匿名变量示例
	// ------------------------------------
	// 注意：匿名变量则通过下划线 _ 来声明，任何赋予它的值都会被丢弃
	_, c1 := GetName()         // 使用匿名变量忽略第一个返回值
	fmt.Printf("c1: %v\n", c1) // 只打印第二个返回值
	fmt.Printf("------------------------------------\n")

	// ------------------------------------
	// 常量定义示例
	// ------------------------------------
	const ( // 通过一个 const 关键字定义多个常量
		PI   float64 = 3.14 // 常量 PI
		zero float32 = 0.0  // 无类型浮点常量
		eof          = -1   // 无类型整型常量
	)
	const d1, d2 float32 = 0, 3    // u = 0.0, v = 3.0，常量的多重赋值
	const d3, d4, d5 = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("zero: %v\n", zero)
	fmt.Printf("eof: %v\n", eof)
	fmt.Printf("d1: %v\n", d1)
	fmt.Printf("d2: %v\n", d2)
	fmt.Printf("d3: %v\n", d3)
	fmt.Printf("d4: %v\n", d4)
	fmt.Printf("d5: %v\n", d5)
	fmt.Printf("------------------------------------\n")

}

func GetName() (nickName string, age int) {
	return "chen", 18
}
