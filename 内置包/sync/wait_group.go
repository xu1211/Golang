package main

import (
	"fmt"
	"sync"
)

// 同步信号量
var wg sync.WaitGroup

func hello() {
	defer wg.Done() //计数器-1
	fmt.Println("Hello Goroutine!")
}
func main() {
	wg.Add(5) //计数器+5

	for i := 0; i < 10; i++ {
		go hello()
	}

	wg.Wait() //等计数器变为0 才执行
	fmt.Println("main goroutine done!")
}
