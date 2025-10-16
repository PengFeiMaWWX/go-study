package main

import "fmt"

func main() {

	var s []int
	printSlice(s)

	// 可以在空切片上追加
	s = append(s, 1)
	printSlice(s)

	//这个切片会按需增长
	s = append(s, 2)
	printSlice(s)

	// 可一次填加多个元素
	s = append(s, 3, 3, 5, 4)
	printSlice(s)
}

func printSlice(s []int) {
	//%v的强大之处在于其通用性，它能智能地适应不同的数据类型：
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
