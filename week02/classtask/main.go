package main

import (
	"fmt"
)

func main() {
	// make 示例
	// 创建并初始化一个切片
	slice := make([]int, 3)
	fmt.Printf("Slice: %v\n", slice)

	// 创建并初始化一个 map
	m := make(map[string]int)
	m["key"] = 1
	fmt.Printf("Map: %v\n", m)

	// 创建并初始化一个通道
	ch := make(chan int)
	fmt.Printf("Channel: %v\n", ch)

	// new 示例
	numPtr := new(int)
	fmt.Printf("Pointer to int: %v, Value: %d\n", numPtr, *numPtr)

	type MyStruct struct {
		a int
		b string
	}
	structPtr := new(MyStruct)
	fmt.Printf("Pointer to struct: %v, a: %d, b: %s\n", structPtr, structPtr.a, structPtr.b)
}
