package main

// Go 官方教程提供的测试包
import (
	"fmt"
	"strings" // Go语言标准库中的一个包，提供了许多用于处理字符串的函数

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)    // 将字符串按一个或多个连续的空格字符（空格、制表符、换行符等）分割
	count := make(map[string]int) // ​make(map[string]int)​：创建空的映射

	for _, word := range words { // _：忽略索引（因为我们不需要索引值） , word 当前单词
		count[word]++ // 统计每个单词出现的次数
	}

	return count
}
func main() {

	testString := "hello world hello go programming go world"
	result := WordCount(testString)
	fmt.Println(result)

	wc.Test(WordCount)

}
