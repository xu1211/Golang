package main

import (
	"fmt"
)

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

	// 圆c1面积：c1.getArea()
	fmt.Println("圆c1的面积 = ", c1.getArea())

	// 通过 普通函数 set() 修改圆半径
	set(&c1, 20)
	fmt.Println("圆c1的面积 = ", c1.getArea())

	// 通过 方法 setRadius() 修改圆半径
	c1.setRadius(30)
	fmt.Println("圆c1的面积 = ", c1.getArea())

}

/*
	定义 结构体 ：
	圆
*/
type Circle struct {
	// 半径
	radius float64
}

/*
	定义 方法 ：
	属于 Circle 类型对象中的方法 getArea()
*/
func (c Circle) getArea() float64 {
	//c.radius 为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

// 定义 函数
func set(c *Circle, radius float64) {
	c.radius = radius
	fmt.Println("圆的半径改为：", c.radius)
}

/*
	定义 方法 ：
	属于 Circle 类型对象中的方法 setRadius(radius float64)
	要更改c的值需要传指针
*/
func (c *Circle) setRadius(radius float64) {
	c.radius = radius
	fmt.Println("圆的半径改为：", c.radius)
}
