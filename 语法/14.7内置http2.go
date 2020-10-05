package main

/*
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
*/
import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {

	// 注册处理方法
	http.Handle("/string", String("I'm a frayed knot."))      //http://localhost:8081/string
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"}) //http://localhost:8081/struct
	// 设置监听的端口
	err2 := http.ListenAndServe("localhost:8081", nil)
	if err2 != nil {
		log.Fatal(err2)
	}
}
