package main

import "fmt"

func main() {
	/*
		一维数组
		定义，赋值，访问
	*/
	fmt.Println("一维数组定义：")
	var array1 [5]int
	fmt.Println(array1)
	println()

	fmt.Println("一维数组定义+赋值：")
	var array2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	println()

	fmt.Println("根据元素的个数来设置数组的大小：")
	var array3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	println()

	fmt.Println("一维数组赋值：")
	var array4 [5]int
	for i := 0; i < len(array4); i++ {
		array4[i] = i + 100
	}
	fmt.Println(array4)
	println()

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
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	// 二维数组 访问
	fmt.Println("二维数组访问：")
	value1 := a1[2][3]
	var value2 int = a2[2][3]
	fmt.Printf("%d %d\n", value1, value2)
	println()

	// 二维数组 遍历
	fmt.Println("二维数组遍历：")
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a3[i][j])
		}
	}
	println()

	println("一维数组长度：", len(a3))
	println("二维数组长度：", len(a3[1]))
}
