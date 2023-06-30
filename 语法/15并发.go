package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

		同一个程序中的所有 goroutine 共享同一个地址空间。

		通过 go 关键字来开启 goroutine, 语法格式：go 函数名( 参数列表 )
	*/
	// 可以开启多个goroutine
	go say("我先 world")
	go say("world")
	say("hello")

	// 主协程退出了，其他任务不执行
	go say("我也 world")
}

// 输出5次
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
