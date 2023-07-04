package main

import (
	"fmt"
	"runtime"
)

/*
runtime.Gosched()
主动让出处理器CPU时间片，允许其他 Goroutine 运行
*/
func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
