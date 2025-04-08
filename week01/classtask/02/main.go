package main

// 判断字符串长度
import (
	"fmt"
)

func countChars(s string) int {
	// 将字符串转换为rune数组
	runes := []rune(s)
	// 返回rune数组的长度
	return len(runes)
}

func main() {
	str := "Hello,世界"
	fmt.Printf("字符串: \"%s\"\n", str)
	fmt.Printf("字符数量: %d\n", countChars(str))

	// 对比直接使用len()的区别
	fmt.Printf("字节长度(len): %d\n", len(str))
}
