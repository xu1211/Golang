package main

import "fmt"

func main() {
	// for 循环
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
		}
		fmt.Println("")
	}

	// for + goto
	for m := 1; m < 10; m++ {
		n := 1
	LOOP:
		if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}
}
