package main

import "fmt"

// Vertex 结构体，存放经纬度
type Vertex struct {
	Lat, Long float64
}

func main() {
	/*
		无序集合，通过 key 来快速检索 value
		定义
			var map_variable map[key_data_type]value_data_type
		初始化
			map_variable := map[key_data_type]value_data_type{key: value}
			map_variable := make(map[key_data_type]value_data_type)
	*/
		fmt.Println("map定义：")
	var map2 map[string]string
	var map1 map[string]Vertex
		fmt.Println("map1：", map1)
		fmt.Println("map2：", map2)
		println()

	// 不初始化 map, 默认是 nil, nil map 不能用来存放键值对，赋值会报错
	//map1["Google"] = Vertex{1.23, 4.56}
	//map2["111"] = "222"

		fmt.Println("map初始化方法1.1 不赋值初始化：")
	map1 = map[string]Vertex{}
		fmt.Println("map1：", map1)

		fmt.Println("map初始化方法1.1 赋值初始化：")
	map1 = map[string]Vertex{
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
		fmt.Println("map1：", map1)

		println()
		fmt.Println("map增加元素：")
	map1["Google"] = Vertex{1.23, 4.56}
		fmt.Println("map1：", map1)
	map1["Alibaba"] = Vertex{7.8, 9.10}
		fmt.Println("map1：", map1)

		println()
	fmt.Println("len 返回的是键值对数目:", len(map1))

		println()
		fmt.Println("map初始化方法2：")
	map2 = make(map[string]string)
	map2["France"] = "巴黎"
	map2["Italy"] = "罗马"
	map2["Japan"] = "东京"
	map2["India "] = "新德里"
		fmt.Println("map2：", map2)

		println()
		fmt.Println("map遍历：")
	for country := range map2 {
		fmt.Println(country, "首都是", map2[country])
	}

		println()
		fmt.Println("map双赋值检测：")
	isHave(map2, "France")

		println()
		fmt.Println("map删除元素：")
	delete(map2, "France")
		fmt.Println("map2：", map2)

		println()
		fmt.Println("map判断元素存在：")
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
