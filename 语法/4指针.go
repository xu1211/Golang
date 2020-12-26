package main

import "fmt"

/*
   var var_name *type	声明指针变量var_name,是指向type类型的指针

   &var_name			取地址,返回存储的内存地址
   *var_name			解引用,返回内存地址对应的值
*/
func main() {
	var a int = 4
	fmt.Println("变量 a = ", a)

	// 声明指针变量，保存的是一个地址值
	var ptr *int
	// 声明指向指针的指针变量
	var pptr **int

	// 指针赋值，取内存地址赋给指针
	ptr = &a
	pptr = &ptr
	fmt.Println("ptr指向a的内存地址 = ", ptr)
	fmt.Println("解引用 *ptr = ", *ptr)
	println("pptr指向ptr的内存地址 = ", pptr)
	fmt.Println("解引用 **pptr = ", **pptr)

	// 空指针，指针被定义后没有分配到任何变量时，它的值为 nil
	var fp *float32
	if fp == nil {
		fmt.Println("fp 为空指针 : ", fp)
	}

	arr := []int{10, 100, 200}
	fmt.Println("数组 arr = ", arr)

	// 声明指针数组
	var arrPtr [3]*int
	// 指针数组赋值
	for i := 0; i < 3; i++ {
		arrPtr[i] = &arr[i]
	}
	fmt.Println("arrPtr = ", arrPtr)
	// 指针数组遍历 解引用
	for i := 0; i < 3; i++ {
		fmt.Println("解引用 *arrPtr[", i, "] = arr[", i, "] = ", *arrPtr[i])
	}

	// 指针作为函数参数
	println("指针作为函数参数:")
	swap(&arr[0], &arr[1])

	fmt.Println("arr[0] = ", arr[0])
	fmt.Println("*arrPtr[0] = ", *arrPtr[0])
	fmt.Println("arr[1] = ", arr[1])
	fmt.Println("*arrPtr[1] = ", *arrPtr[1])
}

func swap(x *int, y *int) {
	fmt.Println("对一个解引用的指针赋值 将会改变这个指针引用的真实地址的值")
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
