package main

import (
	"fmt"
	"time"
)

/*
	select 从多个通道接收值
		每个 case 必须对应一个通道的通信（接收或发送）过程; select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句

		有case 可运行时: 会执行case对应的语句
		没有 case 可运行: select将阻塞 一直循环检测条件，直到有 case 可运行
			此时如果有`default`分支,会执行default对应的语句并退出
		有多个 case 都可以运行: Select 会随机公平地选出一个执行。其他不会执行。
		channel关闭: 关闭后select可以继续读取，直到 channel 中的数据全部读取完毕.
			如果继续读数据，得到的是零值; 可通过 num, ok := <-ch 判断channel是否关闭

		使用 time.After 来设置超时时间

	使用 for select 可以循环读取
*/

// 多case , 超时时间
func selectDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- i
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case num := <-ch1:
			fmt.Println("Received from ch1:", num)
		case num := <-ch2:
			fmt.Println("Received from ch2:", num)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			return
		}
	}
}

// Default分支,关闭channel
func selectDefaultAndCancel() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for {
		select {
		// ch关闭后会一直读出零值, 所以需要判断ch是否关闭
		//case num := <-ch:
		//    fmt.Println("Received from ch:", num)
		case num, ok := <-ch:
			if ok {
				fmt.Println("Received:", num)
			} else {
				fmt.Println("Channel closed!")
				return
			}
		default: // 没读到值就执行default
			fmt.Println("No data received yet.")
			time.Sleep(time.Second)
		}
	}
}
func main() {
	selectDemo()
	fmt.Println("--------------------------------")
	selectDefaultAndCancel()
}
