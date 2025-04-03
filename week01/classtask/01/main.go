package main

import "fmt"

func countChars(s string) int {
	// 将字符串转换为rune数组
	runes := []rune(s)
	// 返回rune数组的长度
	return len(runes)
}

func main() {
	// 定义变量
	name := "小明"   // string类型
	age := 23      // int类型
	gender := true // bool类型，true表示男，false表示女

	// 逐行输出信息
	fmt.Println("姓名:", name)
	fmt.Println("年龄:", age)

	// 根据gender值输出性别
	if gender {
		fmt.Println("性别: 男")
	} else {
		fmt.Println("性别: 女")
	}

}
