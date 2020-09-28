package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		1.函数作为另外一个函数的实参
	*/
	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	// 使用函数
	fmt.Println(getSquareRoot(9))
}
