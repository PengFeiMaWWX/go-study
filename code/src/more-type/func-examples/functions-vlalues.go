package main

import (
	"fmt"
	"math"
)

// 接受一个函数 fn，该函数需要两个 float64参数并返回一个 float64
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func main() {

	// 定义函数名为hypot
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	// 将hypot函数作为参数传给compute
	fmt.Println(compute(hypot))

	fmt.Println(compute(math.Pow))
}

/**
函数值
函数也是值。它们可以像其他值一样传递。

函数值可以用作函数的参数或返回值。
*/
