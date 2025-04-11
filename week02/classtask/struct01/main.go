package main

import (
	"encoding/json"
	"fmt"
)

// Person 定义结构体
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// NewPerson 构造函数
func NewPerson(name string, age int, email string) Person {
	return Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func NewPersonPtr(name string, age int, email string) *Person {
	return &Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

// PrintPerson 打印Person信息
func PrintPerson(p Person) {
	// 打印结构体信息
	fmt.Println("Person信息:")
	fmt.Printf("姓名: %s\n", p.Name)
	fmt.Printf("年龄: %d\n", p.Age)
	fmt.Printf("邮箱: %s\n", p.Email)

	// 转换为JSON并打印
	// jsonData, err := json.Marshal(p)
	jsonData, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("JSON转换错误:", err)
		return
	}
	fmt.Println("JSON格式:")
	fmt.Println(string(jsonData))
}

func main() {
	// 创建Person实例
	person := NewPerson("张三", 25, "zhangsan@example.com")

	// 打印信息
	PrintPerson(person)
}
