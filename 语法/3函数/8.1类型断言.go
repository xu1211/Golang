package main

import (
	"fmt"
	"strconv"
)

/**
空接口可以存储任意类型的值
想要判断空接口中的值这个时候就可以使用类型断言
语法格式： x.(T)
    x：表示类型为interface{}的变量
    T：表示断言x可能是的类型。
*/

type structTypes interface {
	structTypes()
}
type stringStruct struct {
	s string
}

func (*stringStruct) structTypes() {}

type intStruct struct {
	i int
}

func (*intStruct) structTypes() {}

func main() {
	// 测试1: interface{} 值为基础类型
	var x interface{}
	x = "string"

	v, ok := x.(string)
	if ok {
		fmt.Println("x是string: " + v)
	} else {
		fmt.Println("string 类型断言失败")
	}

	i, ok := x.(int)
	if ok {
		fmt.Println("x是int: " + strconv.Itoa(i))
	} else {
		fmt.Println("int 类型断言失败")
	}
	fmt.Println("--------------------------------")
	justifyType(x)
	x = 1
	justifyType(x)
	x = true
	justifyType(x)
	fmt.Println("--------------------------------")
	// 测试2: interface{} 值为struct
	alltype()
}
func alltype() {
	var a structTypes
	a = &stringStruct{s: "s"}
	fmt.Println(a)
	if v, ok := a.(*stringStruct); ok {
		fmt.Println(v.s)
	}

	a = &intStruct{i: 123}
	fmt.Println(a)
	if v, ok := a.(*intStruct); ok {
		fmt.Println(v.i)
	}
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
