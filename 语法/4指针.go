package main

import "fmt"

/*
   &	返回变量存储地址
   *	指针变量
*/
func main() {
	// int变量
	var a int = 4
	// 指针变量，保存的是一个地址值
	var ptr *int

	// 指针变量ptr 指向 a变量存储地址
	ptr = &a
	fmt.Printf("a 为  %d\n", a)
	println("ptr为", ptr)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
