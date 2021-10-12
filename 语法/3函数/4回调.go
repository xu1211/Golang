package main

import "fmt"

/*
go语言中指通过函数指针调用函数:
    把函数的指针（地址）作为参数传递给另一个函数，这个指针被用来调用其所指向的函数
*/

// 声明一个函数类型，用于回调函数的指针
type Callback func(int) int

func main() {
    //方式1
	testCallBack(2, func(x int) int {
		fmt.Printf("我是内部回调函数，x：%d\n", x)
		return x
	})
    //方式2
	testCallBack(1, callBack)
}

// 调用方
func testCallBack(x int, callback Callback) {
    // where 达成条件，执行回调函数
	callback(x)
}

// 实现方 回调函数
func callBack(x int) int {
	fmt.Printf("我是回调函数，x：%d\n", x)
	return x
}
