package main

import "fmt"

// var 变量组，一般用于声明全局变量
var (
	vname1 int
	vname2 string
)

func main() {
	// 变量声明，默认零值
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("默认零值: %v %v %v %q\n", i, f, b, s)

	// 单个变量声明+赋值
	var s1 string = "变量s1赋值"
	fmt.Println(s1)

	// 多个变量声明
	var i1, i2 int
	var s2, s3 string
	fmt.Println(i1, i2, s2, s3)

	// 多个变量赋值
	i1, i2, s2, s3 = 1, 2, "abc", "def"
	fmt.Println(i1, i2, s2, s3)

	// 多个变量声明+赋值
	var s4, s5 string = "变量s4赋值", "变量s5赋值"
	fmt.Println(s4, s5)

	// 根据值自行判定变量类型
	var d = true
	fmt.Println("自行判定变量类型: ", d)

	// 省略 var ，只能在函数体中出现
	v := "NO var" // 等同 var v string = "NO var"
	fmt.Println(v)
}
