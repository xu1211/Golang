package main

import (
	"fmt"
	"io"
	"net/http"
)
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello, world!\n")
	})
	fmt.Println("Please Visit -  http://localhost:9999")
	http.ListenAndServe(":9999",nil)

}