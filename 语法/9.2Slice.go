package main

import "fmt"

func main() {
	/*
		切片
			定义
				1.声明一个未指定大小的数组		如：slice1 - slice3
				2.使用make()函数 make([]T, length, capacity)		如：slice4 - slice6
				4.截取数组或切片src := tar[startIndex:endIndex]		如：slice7 - slice9

			初始化
				切片在未初始化之前默认为nil，长度为0, 空(nil)切片		如：slice1,slice2

			len() 函数：获取长度
			cap() 函数：计算容量
			append() 函数：追加新元素，(切片容量不够时会自动扩容, 扩容规则参考$GOROOT/src/runtime/slice.go源码)
			copy() 函数：拷贝切片

	*/
	println("切片定义：")
	var slice1 []int
	printSlice(slice1)

	slice2 := []int{}
	printSlice(slice2)

	slice3 := []int{1, 2, 3}
	printSlice(slice3)

	var slice4 []int = make([]int, 5)
	printSlice(slice4)

	slice5 := make([]int, 5)
	printSlice(slice5)

	var slice6 = make([]int, 3, 5)
	printSlice(slice6)

	println("切片截取：")
	var arr = [5]int{1, 2, 3, 4, 5}
	slice7 := arr[:]
	printSlice(slice7)

	slice8 := arr[1:4]
	printSlice(slice8)

	slice9 := slice7[1:]
	printSlice(slice9)

	println("切片append()：")
	slice1 = append(slice1, 1)
	printSlice(slice1)

	slice1 = append(slice1, 2, 3, 4, 5)
	printSlice(slice1)

	println("切片copy()：")
	copy(slice3, slice1)
	printSlice(slice3)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
