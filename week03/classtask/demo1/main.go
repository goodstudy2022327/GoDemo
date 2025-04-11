package main

import (
	test "demo1/cal" // 导入自定义包
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(test.Add(1, 2))
	fmt.Println(test.Sub(1, 2))
}
