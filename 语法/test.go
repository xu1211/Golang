package main

import "fmt"

func main() {
	add_func := add(1, 2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())
}

// 闭包使用方法
func add(x1, x2 int) func() (int, int, int) {
	i := 0
	return func() (int, int, int) {
		i++
		return i, x1, x2
	}
}
