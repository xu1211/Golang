package main

import "fmt"

func main() {
	/*
		一维数组
		定义，赋值，访问
	*/
	var array1 [5]int
	list(array1)
	println()

	var array2 = [5]int{1, 2, 3, 4, 5}
	list(array2)
	println()

	var array3 = [...]int{1, 2, 3, 4, 5}
	list(array3)
	println()

	var array4 [5]int
	for i := 0; i < len(array4)-1; i++ {
		array4[i] = i + 100
	}
	list(array4)

	/*
		二维数组
		定义，赋值
	*/
	var a1 = [3][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}

	var a2 = [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}

	var a3 = [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11}}

	// 二维数组 访问
	value1 := a1[2][3]
	var value2 int = a2[2][3]
	fmt.Printf("%d %d\n", value1, value2)

	// 二维数组 遍历
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a3[i][j])
		}
	}
}

func list(inArr [5]int) {
	for j := 0; j < len(inArr)-1; j++ {
		fmt.Printf("Element[%d] = %d\n", j, inArr[j])
	}
}
