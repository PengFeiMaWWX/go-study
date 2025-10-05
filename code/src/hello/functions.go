package main

import "fmt"

/**
函数可接受零个或多个参数。

在本例中，add 接受两个 int 类型的参数。

注意类型在变量名的 后面。
 */
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
