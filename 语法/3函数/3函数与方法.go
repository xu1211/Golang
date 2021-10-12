package main

import (
	"fmt"
	"math"
)
/*
Go 语言中同时有 函数 和 方法 。
方法是一个 作用于 接受者receiver的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
接受者receiver 的概念类似this，self

函数: func function_name( [parameter list] ) [return_types] {函数体}
方法: func (var receiver) method_name( [parameter list] ) [return_types] {方法体}

Go 自动处理方法调用时的值和指针之间的转化。
	1. 方法接受器类型 为值类型: 不能改变结构体的数据
	2. 方法接受器类型 为指针类型: 可以改变结构体的数据，避免在方法调用时产生一个拷贝，
*/

// 接受者receiver
// 定义结构体 Circle：圆
type Circle struct {
	// 半径
	radius float64
}

/*
定义方法：(c Circle) getArea() float64
	接收器类型：Circle 值类型
	传参：无
	返回：float64
*/
func (c Circle) getArea() float64 {
	fmt.Println("c.radius = ", c.radius)
	return 3.14 * c.radius * c.radius
}

func (c Circle) setRadius1(radius float64) {
	c.radius = radius
	fmt.Println("值类型传递 不能改变结构体的数据：", c.radius)
}
/*
定义方法 ：(c *Circle) setRadius(radius float64)
	接收器类型：*Circle 指针类型
	传参：float64
	返回：无
*/
func (c *Circle) setRadius(radius float64) {
	c.radius = radius
	fmt.Println("方法setRadius 圆半径为：", c.radius)
}

// 定义函数
func set(c *Circle, radius float64) {
	c.radius = radius
	fmt.Println("函数set 圆半径为：", c.radius)
}

// 可以对包中的 任意 类型定义任意方法。但是 不能对来自其他包的类型或基础类型定义方法。
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func main() {
	// 圆c1
	var c1 Circle

	// 普通赋值
	c1.radius = 1.00
	fmt.Println("普通赋值 圆半径为：", c1.radius)

	// 通过 Circle类型 的方法getArea() 获取圆面积
	fmt.Println("c1.getArea() = ", c1.getArea())
	println()

	// 通过 普通函数 set() 赋值
	set(&c1, 10)
	fmt.Println("c1.getArea() = ", c1.getArea())
	println()

	// 通过 Circle.setRadius() 赋值
	c1.setRadius(100)
	c1.setRadius1(1)
	fmt.Println("c1.getArea() = ", c1.getArea())
	println()

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
