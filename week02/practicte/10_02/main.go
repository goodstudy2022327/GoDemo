package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// 学生结构体
type Student struct {
	ID    int
	Name  string
	Age   int
	Major string
}

// 全局变量：用于存储学生信息
var students []Student
var dataFile = "students.json"

// 定义学生信息存储，声明并初始化
// var students = make(map[int]Student)

// 加载数据（从文件）
func loadFromFile() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		students = []Student{} // 初始化
		return
	}
	json.Unmarshal(file, &students)
}

// 加载数据（从文件）, 使用bufio缓冲
// func loadFromFile() {
// 	file, err := os.Open(dataFile)
// 	if err != nil {
// 		students = []Student{} // 初始化
// 		return
// 	}
// 	defer file.Close()

// 	// 1. 使用bufio创建Reader
// 	reader := bufio.NewReader(file)
// 	// 2. 创建JSON解码器。传入上面的缓冲读取器reader
// 	decoder := json.NewDecoder(reader)
// 	// 3. 解码并修改结构体的值。传入结构体的地址
// 	err = decoder.Decode(&students)
// 	if err != nil {
// 		fmt.Println("读取数据失败:", err)
// 		students = []Student{}
// 	}
// }

// 保存数据（到文件）
// func saveToFile() {
// 	data, _ := json.MarshalIndent(students, "", "  ")
// 	os.WriteFile(dataFile, data, 0644)
// }

// 保存数据（到文件）, 使用bufio缓冲
func saveToFile() {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer file.Close()

	// 1. 使用bufio创建Writer
	writer := bufio.NewWriter(file)
	// 最后需要刷新缓冲区，写入文件
	defer writer.Flush()

	// 2. 创建JSON编码器。传入上面的缓冲写入器Writer
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ") // 设置缩进，美化输出
	// 3. 结构体的值写入文件
	err = encoder.Encode(students)
	if err != nil {
		fmt.Println("写入数据失败:", err)
	}
}

// 添加学生
func addStudent() {
	var s Student
	fmt.Print("请输入学号: ")
	// 由于传递了变量的地址，fmt.Scanln 可以直接修改 s 结构体的字段值
	fmt.Scanln(&s.ID) // 读取输入并存储到student的ID中

	fmt.Print("请输入姓名: ")
	fmt.Scanln(&s.Name)
	fmt.Print("请输入年龄: ")
	fmt.Scanln(&s.Age)
	fmt.Print("请输入专业: ")
	fmt.Scanln(&s.Major)

	students = append(students, s)
	saveToFile()
	fmt.Println("✅ 添加成功！")
}

// 查看所有学生
func viewStudents() {
	fmt.Println("【学生列表】")
	for _, s := range students {
		fmt.Printf("ID:%d, 姓名:%s, 年龄:%d, 专业:%s\n", s.ID, s.Name, s.Age, s.Major)
	}
}

// 修改学生
func updateStudent() {
	var id int
	fmt.Print("请输入要修改的学号: ")
	fmt.Scanln(&id)
	for i, s := range students {
		if s.ID == id {
			fmt.Print("请输入新姓名: ")
			fmt.Scanln(&students[i].Name)
			fmt.Print("请输入新年龄: ")
			fmt.Scanln(&students[i].Age)
			fmt.Print("请输入新专业: ")
			fmt.Scanln(&students[i].Major)
			saveToFile()
			fmt.Println("✅ 修改成功！")
			return
		}
	}
	fmt.Println("未找到该学生")
}

// 修改学生信息
// func modifyStudent() {
// 	var id int
// 	fmt.Print("输入要修改的学生ID：")
// 	fmt.Scan(&id)

// 	if _, exists := students[id]; !exists {
// 		fmt.Println("未找到该学生信息！")
// 		return
// 	}

// 	fmt.Print("输入新的姓名：")
// 	fmt.Scan(&students[id].Name)
// 	fmt.Print("输入新的年龄：")
// 	fmt.Scan(&students[id].Age)
// 	fmt.Print("输入新的年级：")
// 	fmt.Scan(&students[id].Grade)
// 	fmt.Println("学生信息修改成功！")
// }

// 删除学生
func deleteStudent() {
	var id int
	fmt.Print("请输入要删除的学号: ")
	fmt.Scanln(&id)
	for i, s := range students {
		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			saveToFile()
			fmt.Println("✅ 删除成功！")
			return
		}
	}
	fmt.Println("未找到该学生")
}

// 主菜单
func main() {
	loadFromFile()
	for {
		fmt.Println("\n====== 学生信息管理系统 ======")
		fmt.Println("1. 添加学生")
		fmt.Println("2. 查看所有学生")
		fmt.Println("3. 修改学生")
		fmt.Println("4. 删除学生")
		fmt.Println("5. 退出")
		fmt.Print("请输入操作编号: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			addStudent()
		case "2":
			viewStudents()
		case "3":
			updateStudent()
		case "4":
			deleteStudent()
		case "5":
			fmt.Println("👋 程序已退出")
			return
		default:
			fmt.Println("❌ 无效的选项，请重新输入")
		}
	}
}
