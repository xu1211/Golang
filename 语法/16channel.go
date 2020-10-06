package main

import "fmt"

/*
多线程会引入线程之间的同步问题，在 golang 中可以使用 channel 作为同步的工具。

通道（channel）是用来传递数据的一个数据结构。可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
通道遵循先进先出原则。
操作符 <- 用于指定通道的方向(发送或接收)。如果未指定方向则为双向通道。
默认情况在另一端准备好之前，发送和接收都会阻塞。
*/

func sum(a []int, c chan int) {
	sum := 0
	fmt.Println(a)
	for _, v := range a {
		sum += v
	}
	// 发送端：把 sum 发送到通道 c，发送方会阻塞 直到接收方从通道中接收了值
	c <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)

	go sum(a[:len(a)/2], c) // [1 2 3]
	go sum(a[len(a)/2:], c) // [4 5 6]

	// 接收端：从通道 c 中接收
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
