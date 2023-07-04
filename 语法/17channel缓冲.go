package main

import (
	"fmt"
	"time"
)

/*
无缓冲的通道:
  又称为阻塞通道 同步通道
  通道必须有接收才让发送，否则报错. 且必须一次接收完成。

带缓冲的通道
	又称为 异步通道
	创建: make 提供第二个参数作为缓冲长度来初始化一个 channel缓冲区： make(chan 元素类型, [缓冲大小])
	特性: 允许 发送端发送数据 和 接收端获取数据 处于异步状态， 缓冲区一满发送端就无法再发送数据

通道不带缓冲时：发送方会阻塞 直到接收方从通道中接收了值。
通道带缓冲时：发送方会阻塞 直到发送的值被拷贝到缓冲区内；
如果缓冲区已满：发送方会阻塞 直到某个接收方获取到一个值。
接收方 在缓冲区空的时候会一直阻塞。
*/
func noBuffer() {
	// 创建一个不带缓冲的 channel
	ch := make(chan int)

	// 启动一个 goroutine，向 channel 中发送数据
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Println("Sent:", i)
		}
		close(ch)
	}()

	// 从 channel 中读取数据
	for num := range ch {
		time.Sleep(time.Second)
		fmt.Println("Received:", num)
	}
}
func buffer() {
	// 创建一个带缓冲的 channel，缓冲大小为 3
	ch := make(chan int, 3)

	// 启动一个 goroutine，向 channel 中发送数据
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("Sent:", i)
		}
		close(ch)
	}()

	// 从 channel 中读取数据
	for num := range ch {
		time.Sleep(time.Second)
		fmt.Println("Received:", num)
	}
}
func main() {
	// 收一发一,阻塞
	noBuffer()
	fmt.Println("--------------------------------")
	// 缓冲多大发多大
	buffer()
}
