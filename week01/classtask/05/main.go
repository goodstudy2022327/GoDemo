package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func listFiles(path string, indent int) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		info, err := entry.Info()
		if err != nil {
			return err
		}

		// 打印缩进、文件名、大小和修改时间
		fmt.Printf("%s%s\tSize: %d bytes\tModified: %s\n",
			getIndent(indent),
			entry.Name(),
			info.Size(),
			info.ModTime().Format(time.DateTime))

		if entry.IsDir() {
			// 递归处理子目录
			listFiles(fullPath, indent+1)
		}
	}
	return nil
}

func getIndent(level int) string {
	return strings.Repeat("  ", level)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	root := os.Args[1]
	fmt.Printf("Listing files in: %s\n", root)
	err := listFiles(root, 0)
	if err != nil {
		log.Fatal(err)
	}
}
