package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}

	combinedSlice := append(slice1, slice2...)

	// 1. map去重
	uniqueMap := make(map[int]bool)
	uniqueSlice := []int{}
	// var uniqueSlice []int

	for _, num := range combinedSlice {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqueSlice = append(uniqueSlice, num)
		}
	}

	fmt.Println("uniqueSlice", uniqueSlice)

	// 2. 先排序，再去重
	// 排序
	sort.Ints(combinedSlice)
	// 去重
	j := 0
	for i := 1; i < len(combinedSlice); i++ {
		if combinedSlice[j] != combinedSlice[i] {
			j++
			combinedSlice[j] = combinedSlice[i]
		}
	}
	uniqueSlice2 := combinedSlice[:j+1]

	fmt.Println("uniqueSlice2", uniqueSlice2)
}
