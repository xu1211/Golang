package main

import "fmt"

/*
fmt 包中定义的 Stringer接口。用字符串描述自己的类型。`fmt`还有许多其他包 使用这个来进行输出

type Stringer struct {
    String() string
}
*/

type IPAddr [4]byte

// 实现接口方法 Stringer.String()
func (ip IPAddr) String() string {
	return fmt.Sprintf("改造后的IPAddr类型 String()： %v.%v.%v.%v\n", ip[0], ip[1], ip[2], ip[3])
}

type Person struct {
	Name string
	Age  int
}

// 实现接口方法 Stringer.String()
func (p Person) String() string {
	return fmt.Sprintf("改造后的Person类型 String()： %v (%v years)\n", p.Name, p.Age)
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v", n, a)
	}

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
