package main

import "fmt"

var i, j = 1, 2

// 变量声明可以包含初始值，每个变量对应一个
// 如果提供了初始值，那么类型可以省略
func main() {

	var c, python, java = true, 3.14, "basic"
	fmt.Println(i, j, c, python, java)
}
