package main

import "fmt"

type Vartex struct {
	Lat, Long float64
}

var m = map[string]Vartex{
	"A": Vartex{
		45.12, 1212.12,
	},
	"B": Vartex{
		212.12, 1212.11,
	},
}

func main() {

	fmt.Println(m)
}

/**
映射字面量
映射的字面量和结构体类似，只不过必须有键名。
*/
