package main

import "fmt"

func main() {
	var name string
	var age int
	var height float64

	fmt.Println("请输入个人信息：")

	// 读取姓名
	fmt.Print("姓名: ")
	fmt.Scanf("%s\n", &name)

	// 读取年龄
	fmt.Print("年龄: ")
	fmt.Scanf("%d\n", &age)

	// 读取身高
	fmt.Print("身高(米): ")
	fmt.Scanf("%f\n", &height)

	// 输出收集到的信息
	fmt.Println("\n您输入的信息是：")
	fmt.Printf("姓名: %s\n", name)
	fmt.Printf("年龄: %d岁\n", age)
	fmt.Printf("身高: %.2f米\n", height)
}
