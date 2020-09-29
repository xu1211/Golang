package main

import "fmt"

// 闭包： 定义在一个函数内部的函数，能够读取其他函数内部变量的函数

// 闭包函数func() int，能够读取func getSequence()函数中 i 变量
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 带参数的闭包函数，闭包函数为func(x3, x4 int) (int, int, int) 闭包函数声明中的返回值不用写具体的形参名称
func add(x1, x2 int) func(int, int) (int, int, int) {
	i := 0
	return func(x3, x4 int) (int, int, int) {
		i += 1
		return i, x1 + x2, x3 + x4
	}
}

func main() {
	// 创建函数 nextNumber
	nextNumber := getSequence()
	fmt.Println("nextNumber 变量i")
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// 创建函数 nextNumber1
	nextNumber1 := getSequence()
	fmt.Println("nextNumber1 变量i")
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	add_func := add(1, 2)
	fmt.Println("add_func 变量i ,add_func参数 x1 + x2,闭包函数参数x3 + x4")
	fmt.Println(add_func(1, 1))
	fmt.Println(add_func(2, 2))
	fmt.Println(add_func(3, 3))
}
