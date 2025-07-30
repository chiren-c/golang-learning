package main

import "fmt"

func main() {
	// 基本使用
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	a1 := months[:6] // 上半年
	a2 := months[6:] // 下半年
	fmt.Println(a1, len(a1), cap(a1))
	fmt.Println(a2)
	fmt.Println("------------------")
	for i := 0; i < len(months); i++ {
		fmt.Println("months[", i, "] =", months[i])
	}
	fmt.Println("------------------")
	for k, v := range months {
		fmt.Println("k[", k, "] =", v)
	}
	// 切片复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	fmt.Println("------------------")
	fmt.Println(slice2)
	// 复制 slice1 到 slice 2
	// copy(slice2, slice1) // 只会复制 slice1 的前 3 个元素到 slice2 中 (替换)
	// fmt.Println(slice2)
	// 复制 slice2 到 slice 1
	copy(slice1, slice2) // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置
	fmt.Println(slice1)
	fmt.Println("------------------")
	// 删除元素
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("s1", s1)
	// s1 = s1[2:]
	// fmt.Println("s1", s1)
	s2 := append(s1[:0], s1[3:]...) // 删除开头三个元素
	fmt.Println("s2", s2)
	s3 := append(s1[:1], s1[4:]...) // 删除中间三个元素
	fmt.Println("s3", s3)
	s4 := s1[:copy(s1, s1[3:])] // 删除开头前三个元素
	fmt.Println("s4", s4)
	fmt.Println("------------------")
	// 数据共享
	// 它们的数组指针指向了同一个数组,所以会同步修改
	// s5: [1 2 6 4 5]
	// s6: [2 6]
	s5 := []int{1, 2, 3, 4, 5}
	s6 := s5[1:3]
	s6[1] = 6
	fmt.Println("s5:", s5)
	fmt.Println("s6:", s6)
	fmt.Println("------------------")
	// 解决数据共享同步修改问题
	// 通过 append 创建一个新的切片,从而避免数据共享
	// 没有进行扩容是不会重新分配内存的也就还有数据共享问题
	s7 := s5[1:]
	s7 = append(s7, 0)
	fmt.Println("s5:", s5)
	fmt.Println("s7:", s7)
	s7[2] = 100
	fmt.Println("s5:", s5)
	fmt.Println("s7:", s7)
}
