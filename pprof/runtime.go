package main

import (
	"os"
	"runtime/pprof"
	"test/pprof/data"
)

func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	pprof.StartCPUProfile(f)
	//停止调用
	defer pprof.StopCPUProfile()

	for i := 0; i < 1000*1000*300; i++ {
		data.Add("test")
	}
}
