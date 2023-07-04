package main

import (
	"fmt"
	"runtime"
)

/*
runtime.Goexit()
终止当前 goroutine 的执行; 不会影响其他 goroutine 的执行
  在 Goexit() 调用之前的代码会被正常执行, 包括 defer
  在 Goexit() 调用之后的代码不会被执行, 包括 defer
*/
func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()

			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
	}
}
