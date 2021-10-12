package main

import "fmt"

/*
不带缓冲区的通道在向通道发送值时，必须及时接收，且必须一次接收完成。

make 提供第二个参数作为缓冲长度来初始化一个 channel缓冲区：
带缓冲区的通道允许 发送端发送数据 和 接收端获取数据 处于异步状态， 缓冲区一满发送端就无法再发送数据

通道不带缓冲时：发送方会阻塞 直到接收方从通道中接收了值。
通道带缓冲时：发送方会阻塞 直到发送的值被拷贝到缓冲区内；
如果缓冲区已满：发送方会阻塞 直到某个接收方获取到一个值。
接收方 在缓冲区空的时候会一直阻塞。
*/
func main() {
    //make 一个通道，最多允许缓存 2 个值。
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
