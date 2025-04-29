/*
遍历数据时如何删掉数组中的某一项
*/
package main

import "fmt"

func removeElement(slice []int, elem int) []int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == elem {
			slice = append(slice[:i], slice[i+1:]...) // 删除元素
			i--                                       // 由于元素被删除，索引减一
		}
	}
	return slice
}

func main() {
	data := []int{1, 2, 3, 4, 2, 5}
	result := removeElement(data, 2)
	fmt.Println(result) // 输出: [1 3 4 5]
}
