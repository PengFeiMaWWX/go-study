package main

import (
	"fmt"
)

/**

类型转换
表达式 T(v) 将值 v 转换为类型 T。

一些数值类型的转换：

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
或者，更加简短的形式：

i := 42
f := float64(i)
u := uint(f)
与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。
*/
func main() {

	var i int = 64
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	i1 := 64
	f1 := float64(i1)
	u1 := uint(f1)
	fmt.Println(i1, f1, u1)

}
