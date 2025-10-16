package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 创建一个长度为dy的二维切片
	picture := make([][]uint8, dy)

	// 遍历每一行
	for i := 0; i < dy; i++ {
		// 为每一行分配一个长度为dx的切片
		picture[i] = make([]uint8, dx)
		// 遍历当前行的每一个像素
		for j := 0; j < dx; j++ {
			// 使用公式计算像素灰度值，并转换为uint8类型
			picture[i][j] = uint8((i + j) / 2)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic) // 显示生成的图像
}
