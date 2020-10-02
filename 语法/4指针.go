package main

import "fmt"

/*
   &	返回变量存储地址
   *	指针变量
*/
func main() {
	var a int = 4
	// 指针变量，保存的是一个地址值
	var ptr *int
	// 指向指针的指针
	var pptr **int

	// 空指针，指针被定义后没有分配到任何变量时，它的值为 nil
	var fp *float32
	if ptr == nil {
		fmt.Printf("fp 为 : %x\n", fp)
	}

	// 指针赋值
	ptr = &a
	pptr = &ptr
	fmt.Printf("a 为  %d\n", a)
	println("ptr 为", ptr)
	fmt.Printf("*ptr 为 %d\n", *ptr)
	println("pptr 为", pptr)
	fmt.Printf("**pptr 为 %d\n", **pptr)

	arr := []int{10, 100, 200}
	// 指针数组
	var arrPtr [3]*int
	// 指针数组赋值
	for i := 0; i < 3; i++ {
		arrPtr[i] = &arr[i]
	}
	// 指针数组遍历
	for i := 0; i < 3; i++ {
		fmt.Printf("arr[%d] = %d\n", i, *arrPtr[i])
	}

	// 指针作为函数参数
	println("指针作为函数参数:")
	swap(&arr[0], &arr[1])
	fmt.Printf("arr[0] = %d\n", *arrPtr[0])
	fmt.Printf("arr[1] = %d\n", *arrPtr[1])
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
