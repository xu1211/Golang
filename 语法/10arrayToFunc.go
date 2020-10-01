package main

import "fmt"

func main() {
	var array = []int{1, 2, 3, 4, 5}
	/* 未定义长度的数组只能传给不限制数组长度的函数 */
	setArray(array)
	/* 定义了长度的数组只能传给限制了相同数组长度的函数 */
	var array2 = [5]int{1, 2, 3, 4, 5}
	setArray2(array2)

	var arr = [][]float32{{-1, -2}, {-3, -4}, {-5}}
	prt(arr)
}

func setArray(params []int) {
	fmt.Println("params array length of setArray is : ", len(params))
}

func setArray2(params [5]int) {
	fmt.Println("params array length of setArray2 is : ", len(params))
}

func prt(arr [][]float32) {
	for i := 0; i < 3; i++ {
		println(arr[i][0])
	}
}
