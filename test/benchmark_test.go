package test

import (
	"fmt"
	"testing"
)

func Benchmark_Add(b *testing.B) {
	var n int
	//b.N 表示这个用例需要运行的次数, 会自动递增,把测试时间控制在1s左右
	for i := 0; i < b.N; i++ {
		n++
	}
}

func Benchmark_Alloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}
