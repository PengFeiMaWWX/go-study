package main

import (
	"fmt"
	"math"
)

/*
*
  - 测试导出的包名
    在go中， 如果一个名字的开头是大写的，那么他就是导出的，那么这个名字就可以被其他包访问到。
	在导入一个包时，你只能引用其中已导出的名字。 任何「未导出」的名字在该包外均无法访问。


*/
func main() {

	//const pi = math.pi
	const pi = math.Pi
	fmt.Println(pi)
}
