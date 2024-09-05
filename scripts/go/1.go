package main

import (
	"fmt"
	"reflect"
)

type Element struct {
	key   string
	value int
}

func removeDuplicatesKeepLastEasy(slice []Element) []Element {
	// 创建一个映射来记录元素的最后出现位置
	lastIndex := make(map[string]int)
	for i, v := range slice {
		lastIndex[v.key] = i // 记录每个元素最后出现的索引
	}

	// 创建一个结果切片，按照最后出现的顺序添加元素
	result := make([]Element, 0, len(lastIndex))
	for i := range slice {
		if index, exists := lastIndex[slice[i].key]; exists && i == index {
			result = append(result, slice[i]) // 添加元素
			delete(lastIndex, slice[i].key)   // 删除已添加的元素，确保只添加一次
		}
	}

	return result
}

func removeDuplicatesKeepLast(slice []Element) []Element {
	// 创建一个映射来记录元素的最后出现位置
	lastIndex := make(map[string]int)

	// 创建一个结果切片，按照最后出现的顺序添加元素
	result := make([]Element, 0)
	for i := range slice {
		if index, exists := lastIndex[slice[i].key]; exists {
			result = append(result[:index], result[index+1:]...)
			delete(lastIndex, slice[i].key) // 删除已添加的元素，确保只添加一次
			for key, value := range lastIndex {
				if value > index {
					lastIndex[key] -= 1
				}

			}

		}
		result = append(result, slice[i])
		lastIndex[slice[i].key] = len(result) - 1
	}

	return result
}

func main() {
	s := []Element{
		{"banana", 20},
		{"apple", 1},
		{"apple", 10},
		{"banana", 2},
		{"apple", 3},
		{"orange", 4},
		{"banana", 5},
		{"grape", 6},
		{"grape", 8},
		{"banana", 10},
		{"orange", 5},
		{"orange", 6},
		{"grape", 8},
		{"parrot", 1},
		{"grape", 8},
		{"banana", 888},
		{"apple", 3},
		{"apple", 3},
		{"apple", 4},
		{"parrot", 1},
	}
	result1 := removeDuplicatesKeepLastEasy(s)
	result2 := removeDuplicatesKeepLast(s)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(reflect.DeepEqual(result1, result2))
}
