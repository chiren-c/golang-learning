package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var a [8]byte // 长度为8的数组，每个元素为一个字节
	// var b [8][8]byte // 长度为8的二维数组，每个元素为一个字节
	// var c [8][8][8]byte // 长度为8的三维数组，每个元素为一个字节
	// var d = [3]int{1,2,3} // 声明时初始化,长度为3的数组，元素为1, 2, 3
	// var e = new([3]int) // 通过 new 初始化,创建一个长度为3的数组，元素为0
	a1 := [5]uint8{1, 2}
	fmt.Println(a1, reflect.TypeOf(a1)) // [1 2 0 0 0] [5]uint8
}
