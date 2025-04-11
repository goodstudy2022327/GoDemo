package main

import "fmt"

func findMostFrequentChar(s string) rune {
	// 创建map统计字符出现次数
	charCount := make(map[rune]int)

	// 统计每个字符出现次数
	for _, char := range s {
		charCount[char]++
	}

	// 找出出现次数最多的字符
	var maxChar rune
	maxCount := 0
	for char, count := range charCount {
		if count > maxCount {
			maxChar = char
			maxCount = count
		}
		fmt.Printf("char:%s,count:%d \n", string(char), count)
	}
	return maxChar
}

func main() {
	input := "aabbccddeeff1122334455gggggg"
	result := findMostFrequentChar(input)
	fmt.Printf("字符串: %s\n出现最多的字符是: %c\n", input, result)
}
