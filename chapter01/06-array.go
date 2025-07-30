package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 数组是值类型
	// 在函数中对数组元素进行修改时
	// 并不会影响原来的数组
	// 这种机制带来的另一个负面影响是当数组很大时
	// 值拷贝会降低程序性能。
	// var a [8]byte // 长度为8的数组，每个元素为一个字节
	// var b [8][8]byte // 长度为8的二维数组，每个元素为一个字节
	// var c [8][8][8]byte // 长度为8的三维数组，每个元素为一个字节
	// var d = [3]int{1,2,3} // 声明时初始化,长度为3的数组，元素为1, 2, 3
	// var e = new([3]int) // 通过 new 初始化,创建一个长度为3的数组，元素为0
	a1 := [5]uint8{1, 2}
	fmt.Println(a1, reflect.TypeOf(a1)) // [1 2 0 0 0] [5]uint8
	b1, b2 := a1[0], a1[len(a1)-1]
	fmt.Println("b1,b2", b1, b2)
	a2 := [3]string{"hello", "work"}
	fmt.Println(a2)
	// 遍历数组
	fmt.Println("---------------------")
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println("i", i, arr[i])
	}
	fmt.Println("---------------------")
	for k, v := range arr {
		fmt.Println("k", k, "v", v)
	}
	fmt.Println("---------------------")
	// 通过二维数组打印九九乘法表
	var multi [9][9]string
	for i := 0; i < 9; i++ {
		for j := 0; j < i; j++ {
			n1 := i + 1  // 2
			n2 := j + 1  // 1
			if n1 < n2 { // 摒除重复的记录
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
		}
	}
	// 打印九九乘法表
	for _, m1 := range multi {
		for _, m2 := range m1 {
			fmt.Printf("%-8s", m2) // 位宽为8，左对齐
		}
		fmt.Println()
	}
}
