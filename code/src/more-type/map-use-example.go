package main

import "fmt"

func main() {

	// 声明并初始化映射
	m := make(map[string]int)

	// 添加键值对
	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3

	// 访问值
	fmt.Println(m["apple"])

	// 检查键是否存在
	/**
	​第一个返回值​：键对应的值（如果键不存在，返回该类型的零值）

	​第二个返回值​：布尔值，表示键是否存在
	*/
	if v, ok := m["banana"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found banana")
	}

	// 删除键
	delete(m, "cherry")
	fmt.Println(m)
}

/**
1
2
map[apple:1 banana:2]
*/
