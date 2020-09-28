package main

import "fmt"

// 声明一个函数类型
type cb func(int) int

func main() {
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
	testCallBack(1, callBack)
}

func testCallBack(x int, f cb) {
	f(x)
}

// 回调函数
func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}
