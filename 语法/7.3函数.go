package main

import (
	"fmt"
	"math"
)

/*
	定义 结构体 ：圆
*/
type Circle struct {
	// 半径
	radius float64
}

/*
	定义 Circle 类型 的方法： getArea()，返回 float64
*/
func (c Circle) getArea() float64 {
	//c.radius 为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

/*
	定义 Circle 指针类型 的方法 ：setRadius(radius float64)，无返回值
*/
func (c *Circle) setRadius(radius float64) {
	c.radius = radius
	fmt.Println("圆的半径改为：", c.radius)
}

// 定义 函数
func set(c *Circle, radius float64) {
	c.radius = radius
	fmt.Println("圆的半径改为：", c.radius)
}

// 可以对包中的 任意 类型定义任意方法。但是不能对来自其他包的类型或基础类型定义方法。
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
Go 语言中同时有 函数 和 方法 。
方法是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
*/
func main() {
	// 圆c1
	var c1 Circle

	// 圆c1半径：c1.radius
	c1.radius = 10.00
	fmt.Println("圆的半径改为：", c1.radius)

	// 通过 Circle类型 的方法getArea() 获取圆面积
	fmt.Println("圆c1的面积 = ", c1.getArea())

	// 通过 普通函数 set() 修改圆半径
	set(&c1, 20)
	fmt.Println("圆c1的面积 = ", c1.getArea())

	// 通过 Circle类型 的方法 setRadius() 修改圆半径
	c1.setRadius(30)
	fmt.Println("圆c1的面积 = ", c1.getArea())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
