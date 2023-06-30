package main

import (
	"sync"
)

// 安全的一次性初始化, 只执行一次
var loadOnce sync.Once

func main() {
	var icons map[string]string
	loadOnce.Do(func() {
		icons = make(map[string]string)
		icons["spades.png"] = "./1.png"
		icons["hearts.png"] = "./2.png"
		icons["diamonds.png"] = "./3.png"
		icons["clubs.png"] = "./4.png"
	})

	for k, v := range icons {
		println(k, v)
	}
}
