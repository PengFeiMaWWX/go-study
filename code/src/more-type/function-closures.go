package main

import "fmt"

/*
*
在adder函数中，我们定义了一个局部变量sum，然后返回一个匿名函数。
这个匿名函数引用了adder函数内部的变量sum。
即使adder函数已经执行完毕，返回的匿名函数仍然可以访问和修改sum变量。
这就是闭包的特性：函数可以访问其词法作用域之外的变量。
*/
func addr() func(int) int { // 函数名 addr()，
	// unc adder()：定义一个名为 adder的函数
	// func(int) int表示：
	// 	- 这是一个函数类型
	//  - 接受一个 int类型的参数
	//  - 返回一个 int类型的值
	sum := 0                 // 定义了一个局部变量sum（初始值为0）。
	return func(x int) int { // 返回一个匿名函数，该匿名函数接受一个int参数x，并将其加到sum上，然后返回当前的sum值。
		sum += x
		return sum
	}
}

func main() {

	pos, neg := addr(), addr() //创建了两个闭包：pos和neg。它们分别是通过调用adder得到的

	for i := 0; i < 10; i++ {

		fmt.Println(
			pos(i),
			neg(-2*i))
	}
}
