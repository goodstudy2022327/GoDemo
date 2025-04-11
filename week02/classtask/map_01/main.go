package main

import "fmt"

func main() {
	// 创建并初始化成绩表map
	scoreMap := map[string]int{
		"小明": 60,
		"小王": 70,
		"张三": 95,
		"李四": 98,
		"王五": 100,
		"张伟": 88,
	}

	// 打印成绩表
	fmt.Println("宿舍成员数学成绩表：")
	for name, score := range scoreMap {
		fmt.Printf("%s: %d\n", name, score)
	}
}
