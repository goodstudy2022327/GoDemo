package main

import "fmt"

// 定义Book结构体
type Book struct {
	// 在这里定义字段
	Title  string
	Author string
	Year   int
}

// 实现FindBooksByAuthor函数
func FindBooksByAuthor(author string, books []Book) []Book {
	// 在这里实现函数逻辑
	var result []Book // 结构体数组切片，只有这样才能append

	for _, book := range books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}

func main() {
	// 结构体切片
	books := []Book{
		{Title: "Go语言编程", Author: "张三", Year: 2020},
		{Title: "算法导论", Author: "李四", Year: 2001},
		{Title: "Go实战", Author: "张三", Year: 2022},
		{Title: "数据库系统", Author: "王五", Year: 2019},
		{Title: "Go Web编程", Author: "张三", Year: 2021},
	}

	// 查找张三的所有书籍
	author := "张三"
	zhangsBooks := FindBooksByAuthor(author, books)

	// 打印结果
	fmt.Printf("作者「%s」的书籍:\n", author)
	for i, book := range zhangsBooks {
		fmt.Printf("%d. 《%s》 (%d年)\n", i+1, book.Title, book.Year)
	}
}
