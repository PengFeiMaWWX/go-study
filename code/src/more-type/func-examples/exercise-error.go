package main

import "fmt"

func Sqrt(f float64) (float64, error) {

	// 当输入为负数时，返回自定义错误
	if f < 0 {
		return 0, ErrorNativeSqrt(f)
	}

	// 使用牛顿迭代法计算平方根[3,4](@ref)
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-f)/(2*z)
	}
	return z, nil
}

// 自定义错误类型
type ErrorNativeSqrt float64

// 实现error 接口的 Errror 方法
func (e ErrorNativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func main() {
	// 测试正数
	if result, err := Sqrt(2); err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("Sqrt(2) = %v\n", result)
	}

	// 测试负数
	if result, err := Sqrt(-2); err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("Sqrt(-2) = %v\n", result)
	}
}
