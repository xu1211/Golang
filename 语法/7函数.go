package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	// 函数调用
	fmt.Printf("最大值是 : %d\n", max(a, b))

	swap1(a, b)
	fmt.Printf("值传递后a b值为 : %d %d\n", a, b)
	swap2(&a, &b)
	fmt.Printf("引用传递后a b值为 : %d %d\n", a, b)
}

// 函数定义
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 值传递
func swap1(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

// 引用传递
func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
