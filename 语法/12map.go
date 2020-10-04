package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	// 不初始化 map 默认是 nil, nil map 不能用来存放键值对
	var map1 map[string]Vertex
	var map2 map[string]string

	/*
		2种方式定义map，
		var map_variable map[key_data_type]value_data_type
		map_variable := make(map[key_data_type]value_data_type)
	*/
	// 1.new
	map1 = map[string]Vertex{
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	map1["Google"] = Vertex{1.23, 4.56}
	map1["Alibaba"] = Vertex{7.8, 9.0}
	fmt.Println(map1)

	// 2.使用 make 函数创建
	map2 = make(map[string]string)

	/* map插入key - value对 */
	map2["France"] = "巴黎"
	map2["Italy"] = "罗马"
	map2["Japan"] = "东京"
	map2["India "] = "新德里"

	// 遍历
	for country := range map2 {
		fmt.Println(country, "首都是", map2[country])
	}

	println()
	isHave(map2, "France")

	/*删除元素*/
	delete(map2, "France")
	isHave(map2, "France")
}

/* 通过双赋值检测某个键存在 */
func isHave(m map[string]string, s string) (string, bool) {
	/*
		flag=true 则键存在,值返回在elem中；
		flag=false 则不存在,elem是map的元素类型的零值
	*/
	elem, flag := m[s]
	if flag {
		fmt.Println(s, "存在，值是", elem)
	} else {
		fmt.Println(s, "不存在")
	}
	return s, flag
}
