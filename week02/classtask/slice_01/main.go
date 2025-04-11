package main

import "fmt"

func main() {
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}

	fmt.Println("初始切片:", numbers)

	// 2. 使用切片操作获取第3到第7个元素（包含第7个）
	mySlice := numbers[2:7] // 左闭右开
	fmt.Println("第3-7个元素:", mySlice)

	// 3. append添加元素
	numbers = append(numbers, 11, 12, 13)
	fmt.Println("添加元素后:", numbers)

	// 4. 删除切片第5个元素
	indexToDelete := 4
	numbers = append(numbers[:indexToDelete], numbers[indexToDelete+1:]...)
	fmt.Println("删除第5个元素:", numbers)

	// 5. 将切片中元素乘2
	for i := range numbers {
		numbers[i] *= 2
	}

	// 6. 打印切片内容和容量
	fmt.Printf("最终切片: %v, 长度: %d, 容量: %d\n", numbers, len(numbers), cap(numbers))

}
