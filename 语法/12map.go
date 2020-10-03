package main

import "fmt"

func main() {
	/*
		2种方式定义map，不初始化 map 默认是 nil, nil map 不能用来存放键值对
		var map_variable map[key_data_type]value_data_type
		map_variable := make(map[key_data_type]value_data_type)
	*/
	var map1 map[string]string
	map1 = map[string]string{"France": "巴黎", "Italy": "罗马"}
	map1["France"] = "巴黎"

	/* 使用 make 函数创建 */
	var map2 map[string]string
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

/*查看元素在集合中是否存在 */
func isHave(m map[string]string, s string) (string, bool) {
	/* flag=true则键存在,值返回在str中；否则不存在 */
	str, flag := m[s]
	if flag {
		fmt.Println(s, "存在，值是", str)
	} else {
		fmt.Println(s，"不存在")
	}
	return s, flag
}
