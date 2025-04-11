package main

import "fmt"

// 在这里实现Swap函数
func Swap(a, b *int) {
	// 你的代码
}

func main() {
	num1 := 5
	num2 := 10
	fmt.Printf("交换前: num1 = %d, num2 = %d\n", num1, num2)
	Swap(&num1, &num2)
	fmt.Printf("交换后: num1 = %d, num2 = %d\n", num1, num2)
}
