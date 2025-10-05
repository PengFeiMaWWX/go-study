package main

import "fmt"

/*
*

多返回值
函数可以返回任意数量的返回值。

swap 函数返回了两个字符串。
:= 是短变量声明操作符，它用于在函数内部同时完成变量的声明和初始化。

string ：Go 的类型名小写，暗示其为基础值类型
*/
func swap(x, y string) (string, string) {
	return y, x
}

func main() {

	a, b := swap("hello", "world")
	fmt.Println(a, b)

}
