package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
对象中的属性首字母必须是大写，才能将结构体对象转换为JSON
变量打上了json标签 `json:"name"` ，表示转化成的json key就用该标签“name”，否则取变量名作为key
`json:"-"` 表示忽略该字段
*/
type Person struct {
	Name  string
	age   int
	Class string `json:"class"`
	Time  int64  `json:"-"`
}

func main() {
	person := Person{"小明", 18, "1班", time.Now().Unix()}
	// //json.Marshal将对象转换为json字符串
	if result, err := json.Marshal(&person); err == nil {
		fmt.Println(string(result))
	}
}
