package main

import "unsafe"

// const 常量组,定义常量组时如果不提供初始值，将使用上行的表达式。
const (
    a = 1
    b
    c
)

// 常量表达式中，函数必须是内置函数
const (
	a1 = "abc"
    b1 = len(a1)
    c1 = unsafe.Sizeof(a1)
)

func main() {
	println(a, b, c)
	println(a1, b1, c1)

	// 单个常量赋值
	const LENGTH int = 10
	const WIDTH int = 5  
	var area int
	area = LENGTH * WIDTH
	println(area)

	// 多重赋值
	const d, e, f = 1, false, "str"
	println(d, e, f)

	// 显式类型定义
	const s1 string = "abc"

	// 隐式类型定义
	const s2 = "abc"

	// 特殊常量iota，第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
	const (
		i1 = iota
		i2
		i3
	)
	println(i1, i2, i3)

	// 同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始。
	const (
		j1 = iota   //0
		j2          //1
		j3          //2
		j4 = "ha"   //"ha"   iota += 1
		j5          //"ha"   iota += 1
		j6 = 100    //100    iota += 1
		j7          //100    iota += 1
		j8 = iota   //7		恢复计数
		j9          //8
	)
	println(j1, j2, j3, j4, j5, j6, j7, j8, j9)
}