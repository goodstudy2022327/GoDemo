// 回文数：给你一个整数x，如果x是一个回文整数，返回true；否则，返回false。回文数是指正
// 序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121是回文，而123不是。

// 1. 使用字符串转换的方法来判断一个整数是否是回文数。
// 把整数转成字符串，然后比较字符串的第一位与最后一位、第二位与倒数第二位。。。以此类推，只要有不一样的就返回false
// 2. 使用数学反转的方法来判断一个整数是否是回文数。
// 通过对整数进行反转来判断是否是回文数。判断反转后的数字与原来的数字是否相同

package main

import "strconv"

func isPalindrome(n int) bool {
	// 将数字转换为字符串
	s := strconv.Itoa(n)
	// 判断字符串是否等于它的反转
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	// 负数不可能是回文数，因为它们有一个负号
	if x < 0 {
		return false
	}

	// 保存原始数字，用于后面的比较
	originNumber := x

	// 用于存储反转后的数字
	reverseNumber := 0

	// 循环提取每一位数字，构建反转数字
	for x > 0 {
		// 取出数字的最后一位，并将其加到反转数字的末尾
		reverseNumber = reverseNumber*10 + x%10
		// 去掉最后一位数字，继续处理剩余的数字
		x = int(x / 10)
	}

	// 比较原始数字与反转后的数字是否相等
	return originNumber == reverseNumber
}
