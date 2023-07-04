package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
runtime.GOMAXPROCS(n)
  设置最大执行线程数, 即同时执行的 goroutine 数量。
  默认情况下，Go 程序会使用所有可用的 CPU 核心来执行 goroutine。但是，在某些情况下，我们可能需要限制 goroutine 的并发度，以避免过度占用系统资源
*/

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	// 获取当前系统的 CPU 核心数
	numCPU := runtime.NumCPU()
	fmt.Println("当前系统的 CPU 核心数:", numCPU, ",但我要限制只用1个")

	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}
