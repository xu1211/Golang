package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"test/pprof/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("test"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}

//http://127.0.0.1:6060/debug/pprof/
