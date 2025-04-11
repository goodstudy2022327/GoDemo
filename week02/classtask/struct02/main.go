package main

import (
	"encoding/json"
	"fmt"
)

// Person 定义结构体，与JSON结构对应
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// JSON格式字符串
	jsonStr := `{"name":"Jane Smith","age":25,"email":"janesmith@example.com"}`

	// 创建Person实例用于接收反序列化结果
	var person Person

	// 将JSON字符串反序列化到结构体
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Println("JSON反序列化失败:", err)
		return
	}

	// 打印
	// 打印结构体字段
	fmt.Println("反序列化结果:")
	fmt.Printf("姓名: %s\n", person.Name)
	fmt.Printf("年龄: %d\n", person.Age)
	fmt.Printf("邮箱: %s\n", person.Email)

}
