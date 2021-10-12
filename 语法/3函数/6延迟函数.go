package main

import "fmt"

func main() {
	/* defer延迟函数：
	原理：deger函数不会立刻调用执行，而是被压入一个栈中。
	     函数返回时 才会按照后进先出的顺序调用被延迟的函数调用。
	*/

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
