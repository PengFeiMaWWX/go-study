package main

import "fmt"

/*
*	初始化语句和后置语句是可选的。封号也可以省略
 */
func main() {

	sum := 1
	for sum < 12 {
		sum += sum
	}

	fmt.Println(sum)
}
