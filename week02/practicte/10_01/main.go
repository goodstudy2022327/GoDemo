package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// 结构体中的ID用不上，根据数据索引来完成
type Task struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

const taskFile = "tasks.json"

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()
	// 使用json.Decoder处理大型的JSON文件，转为结构体切片
	err = json.NewDecoder(file).Decode(&tasks)
	// 使用json.Encoder流式写入
	return tasks, err
}

func saveTasks(tasks []Task) error {
	// 结构体数组切片进行序列化为json
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	// 写入文件
	return os.WriteFile(taskFile, data, 0644)
}

// 添加
func addTask(content string) {
	tasks, _ := loadTasks()
	tasks = append(tasks, Task{Content: content, Done: false})
	saveTasks(tasks)
	fmt.Println("任务已添加:", content)
}

// 列出任务
func listTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	for i, task := range tasks {
		if !task.Done {
			fmt.Printf("%d. %s\n", i+1, task.Content)
		}
	}
}

// 完成任务
func completeTask(index int) {
	tasks, err := loadTasks()
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("任务编号错误")
		return
	}
	tasks[index-1].Done = true
	saveTasks(tasks)
	fmt.Println("任务已完成:", tasks[index-1].Content)
}

// 删除任务
func deleteTask(index int) {
	tasks, err := loadTasks()
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("任务编号错误")
		return
	}
	fmt.Println("任务已删除:", tasks[index-1].Content)
	tasks = append(tasks[:index-1], tasks[index:]...)
	saveTasks(tasks)
}

func main() {
	// 定义命令行参数
	// 添加任务命令
	addCmd := flag.NewFlagSet("add", flag.ExitOnError) // 参数名称
	addContent := addCmd.String("content", "", "任务内容") // 参数内容
	// addCmd := flag.String("add", "", "任务内容")

	// 列出任务命令 - 不需要参数
	// listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// 完成任务命令
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	doneIndex := doneCmd.Int("id", 0, "要标记为完成的任务ID")

	// 删除任务命令
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteIndex := deleteCmd.Int("id", 0, "要删除的任务ID")

	// os.Args获取命令行参数的一个字符串切片（[]string）
	if len(os.Args) < 2 {
		fmt.Println("使用方法: add|list|done|delete")
		return
	}

	switch os.Args[1] {
	case "add":
		// addCmd 是一个 *flag.FlagSet 对象，代表 add 子命令的参数解析器
		// Parse() 方法用于解析传入的参数
		addCmd.Parse(os.Args[2:]) // 获取addContent的值
		if *addContent != "" {
			addTask(*addContent)
		} else {
			fmt.Println("请输入任务内容")
		}
	case "list":
		listTasks()
	case "done":
		doneCmd.Parse(os.Args[2:])
		if *doneIndex > 0 {
			completeTask(*doneIndex)
		} else {
			fmt.Println("请输入任务编号")
		}
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteIndex > 0 {
			deleteTask(*deleteIndex)
		} else {
			fmt.Println("请输入任务编号")
		}
	default:
		fmt.Println("未知命令")
	}
}

// go run main.go add -content="完成作业"
// go run main.go list
// go run main.go done -id=1
// go run main.go delete -id=1
