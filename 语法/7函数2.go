package main

import "fmt"

// 闭包使用方法，函数声明中的返回值(闭包函数)不用写具体的形参名称
func add(x1, x2 int) func(int, int) (int, int, int) {
	i := 0
	return func(x3, x4 int) (int, int, int) {
		i += 1
		return i, x1 + x2, x3 + x4
	}
}

func main() {
	add_func := add(1, 2)
	fmt.Println(add_func(4, 5))
	fmt.Println(add_func(1, 3))
	fmt.Println(add_func(2, 2))
}
