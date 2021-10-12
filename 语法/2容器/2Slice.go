package main

import "fmt"

func main() {
	/*
		切片 slice
			底层是数组指针,数据结构如下:
				struct Slice{
					byte*    array;       // 实际数据
					uintgo    len;        // 长度
					uintgo    cap;        // 容量
				};
			- 切片拥有 长度 和 容量。
				- 长度 是它所包含的元素个数。
				- 容量 是从它的第一个元素开始数，到其底层数组元素末尾的个数。
			创建:
				1.声明一个未指定大小的数组		如：slice1，slice2，slice3
				2.使用make()函数 make([]T, length, capacity)		如：slice4 - slice6
				3.截取数组或切片src := tar[startIndex:endIndex]		如：slice7 - slice9

			初始化:
				切片在未初始化之前默认为nil，长度为0, 空(nil)切片		如：slice1,slice2

			len() 函数：获取长度
			cap() 函数：计算容量
			append() 函数：追加新元素，(切片容量不够时会自动扩容, 扩容规则参考$GOROOT/src/runtime/slice.go源码)
			copy() 函数：拷贝切片

	*/
	fmt.Println("切片创建方法1：")
	var slice1 []int
	fmt.Print("slice1：")
	printSlice(slice1)

	slice2 := []int{}
	fmt.Print("slice2：")
	printSlice(slice2)

	slice3 := []int{1, 2, 3}
	fmt.Print("slice3：")
	printSlice(slice3)

	println()
	fmt.Println("切片创建方法2：")
	var slice4 []int = make([]int, 5)
	fmt.Print("slice4：")
	printSlice(slice4)

	slice5 := make([]int, 5)
	fmt.Print("slice5：")
	printSlice(slice5)

	var slice6 = make([]int, 3, 5)
	fmt.Print("slice6：")
	printSlice(slice6)

	println()
	fmt.Println("切片创建方法3：")
	var arr = [5]int{1, 2, 3, 4, 5}
	slice7 := arr[:]
	fmt.Print("slice7：")
	printSlice(slice7)

	slice8 := arr[1:4]
	fmt.Print("slice8：")
	printSlice(slice8)

	slice9 := slice7[1:]
	fmt.Print("slice9：")
	printSlice(slice9)

	slice10 := slice7[:1]
	fmt.Print("slice10：")
	printSlice(slice10)

	println()
	println("slice10 append()追加1个元素：")
	slice10 = append(slice10, 2)
	fmt.Print("slice10：")
	printSlice(slice10)

	fmt.Print("slice10 append()追加多个元素：")
	slice10 = append(slice10, 3, 4, 5, 6)
	fmt.Print("slice10：")
	printSlice(slice10)

	println("slice10 append()追加切片：")
	slice10 = append(slice10, []int{1,2,3}...)
	fmt.Print("slice10：")
	printSlice(slice10)

	println()
	println("切片9 copy 10：")
	copy(slice9, slice10)
	fmt.Print("slice9：")
	printSlice(slice9)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
