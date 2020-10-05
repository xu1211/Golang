package main

/*
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

*/
import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello golang http!")
}

func main() {
	var h Hello
	// 设置监听的端口
	err := http.ListenAndServe("localhost:8080", h)
	if err != nil {
		log.Fatal(err)
	}
}
