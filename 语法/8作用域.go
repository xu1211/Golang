package main

import "fmt"

/*
局部变量：函数内定义的变量
全局变量：函数外定义的变量
形式参数：函数定义中的变量
*/

/* 1.声明全局变量 */
var g int
var i int = 1

func main() {
	/* 2.声明main 函数 局部变量 */
	var a, b int
	a = 10
	b = 20
	g = a + b
	fmt.Printf("局部变量：a = %d, b = %d ，全局变量 g = %d\n", a, b, g)

	// 全局变量与局部变量名称可以相同，但是 局部变量会被优先考虑
	var i int = 2
	fmt.Printf("变量i = %d\n", i)

	/* 3.形式参数 */
	sum(a, b)

	// 通过花括号来控制变量的作用域，花括号中的变量是单独的作用域，同名变量会覆盖外层
	{
		a := 3
		fmt.Println("in a = ", a)
	}
	fmt.Println("out a = ", a)
}

// sum
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	return a + b
}
