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

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:8080", h)
	if err != nil {
		log.Fatal(err)
	}
}
