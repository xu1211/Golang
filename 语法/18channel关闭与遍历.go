package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	/*
		channel关闭：
			只有发送者才能 close 关闭一个 channel 来表示再没有值会被发送了。
			通常情况下无需关闭它们。只有在需要告诉接收者没有更多的数据的时候才有必要进行关闭，例如中断一个 `range`
			关闭通道并不会丢失里面的数据，只是让读取通道数据的时候不会读完之后一直阻塞等待新数据写入
	*/
	close(c)

	//向一个已经关闭的 channel 发送数据会引起 panic
	//c <- 1
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	/*
		接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭
		通道接收不到数据后 ok 就为 false
	*/
	// v, ok := <-c

	/*
		channel遍历
		   range 函数遍历每个从通道接收到的数据，不断从 channel 接收值，直到channel被关闭。
		   如果channel通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候阻塞。
	*/
	for i := range c {
		fmt.Println(i)
	}
}
