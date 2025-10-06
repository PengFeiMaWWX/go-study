package main

import "fmt"

// 下面是一个计算矩形面积和周长的例子：
// 返回值部分 (area float64, perimeter float64)明确地为两个返回值分别命名为 area（面积）和 perimeter（周长）
func rectProps(length, width float64) (area float64, perimeter float64) {

	area = length * width
	perimeter = 2 * (length + width)
	return // 裸返回，自动返回 area 和 perimeter 的当前值
}

func main() {
	var area, perimeter float64           // 预先声明变量
	area, perimeter = rectProps(5.0, 3.0) // 然后赋值
	fmt.Println("矩形的面积是：", area)
	fmt.Println("矩形的周长是：", perimeter)
}
