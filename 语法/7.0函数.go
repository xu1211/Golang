package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	// 函数调用
	fmt.Println("a,b最大值是 : ", max(a, b))

	fmt.Println(swapInt(a, b))
	fmt.Printf("值传递后a b值为 : %d %d\n", a, b)

	swapIntRef(&a, &b)
	fmt.Printf("引用传递后a b值为 : %d %d\n", a, b)

	sumAll(a)
	sumAll(a, b)

	fmt.Println(split(a))
}

/* 函数定义
输入： 两个int num1, num2
输出： 一个int
*/
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

/*
值传递：x int, y int 缩写为 x, y int
多值返回：返回值类型写在括号里
*/
func swapInt(x, y int) (int, int) {
	return y, x
}

// 引用传递
func swapIntRef(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

/*
可变参数函数：
	使用任意数目的值作为参数。
*/
func sumAll(nums ...int) {
	fmt.Print("变参：", nums, " sum = ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


/*
命名返回值：
  返回值可以被命名，并且像变量那样使用
  没有参数的 return 语句返回结果的当前值
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
