package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}
func main() {
	c := make(chan int, 10)
	fibonacci(cap(c), c)
	/*
		channel遍历
		   range 函数遍历每个从通道接收到的数据，不断从 channel 接收值，直到channel被关闭。
		   如果channel通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候阻塞。
	*/
	for i := range c {
		fmt.Println(i)
	}
}
