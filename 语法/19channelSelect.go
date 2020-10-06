package main

import (
	"fmt"
)

/*
	select
		每个 case 必须是一个通信操作，发送/接收。
		如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
		如果没有 case 可运行，select将阻塞，循环检测条件，直到有 case 可运行，则执行并退出，否则一直循环检测。
		如果没有 case 可运行，`default` 分支会被执行。为了非阻塞的发送或者接收，可使用 default 分支
*/
func chann(c1, c2, c3 chan int) {
	c1 <- 123
	close(c2)
}

func main() {
	// 通道c1, c2, c3
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go chann(c1, c2, c3)

	//var i1 int
	select {
	case i1 := <-c1:
		fmt.Printf("从c1接收到： %d \n", i1)
	case i2, ok := (<-c2): // 等同: i2, ok := <-c2
		if ok {
			fmt.Printf("从c2接收到： %d \n", i2)
		} else {
			fmt.Printf("c2 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
