package main

import "fmt"

func main() {
	var ch1 chan int // 声明一个传递整型的通道
	// var ch2 chan bool    // 声明一个传递布尔型的通道
	// var ch3 chan []int   // 声明一个传递int切片的通道
	// fmt.Println("ch1:", ch1, " ch2:", ch2, "ch3:", ch3)
	ch1 = make(chan int) // 初始化通道,无缓冲通道
	go func(ch chan int) {
		ch <- 10 // 把10发送到ch中,无缓冲通道上的发送操作会阻塞
	}(ch1)

	x := <-ch1 // 从ch中接收值并赋值给变量x
	// <-ch1      // 从ch中接收值，忽略结果
	fmt.Println("x:", x)
	close(ch1)
	fmt.Println("----------------------")
	// 有缓冲通道
	ch1 = make(chan int, 1) // 初始化通道,有缓冲通道
	l := len(ch1)           // 获取通道的长度
	fmt.Println("缓冲通道长度 b:", l)
	ch1 <- 10
	l = len(ch1) // 获取通道的长度
	fmt.Println("缓冲通道长度 a:", l)
	fmt.Println("发送成功")
	fmt.Println("----------------------")
	cha1 := make(chan int, 2)
	cha2 := make(chan int, 2)
	go func() {
		for i := 1; i < 10; i++ {
			cha1 <- i
		}
		close(cha1) // 关闭通道
	}()
	go func() {
		for {
			i, ok := <-cha1 // 从cha1中接收值
			if !ok {
				break // 如果通道已关闭，退出循环
			}
			cha2 <- i * 2 // 将接收到的值乘以2后发送到cha2
		}
		close(cha2) // 关闭cha2通道
	}()
	for i := range cha2 { // 从cha2中接收值
		fmt.Println(i) // 打印接收到的值
	}
	fmt.Println("---------单向通道-------------")
	chan1 := make(chan int)
	chan2 := make(chan int)
	go counter1(chan1)
	go squarer(chan2, chan1)
	printer(chan2)
}

// chan1<- int是一个只能发送的通道，可以发送但是不能接收；
func counter1(chan1 chan<- int) {
	for i := 0; i < 10; i++ {
		chan1 <- i
	}
	close(chan1)
}

// <-chan1 int是一个只能接收的通道，可以接收但是不能发送。
func squarer(chan2 chan<- int, chan1 <-chan int) {
	for i := range chan1 {
		chan2 <- i * i
	}
	close(chan2)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println("printer", i)
	}
}
