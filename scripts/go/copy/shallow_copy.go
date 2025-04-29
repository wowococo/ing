// 
package main

import "fmt"

func main() {
    // 定义一个切片
    original := []int{1, 2, 3}
    // 浅拷贝，只复制引用
    copySlice := original

    // 修改拷贝切片的元素
    copySlice[0] = 100

    // 输出原始切片和拷贝切片
    fmt.Println("Original slice:", original)
    fmt.Println("Copied slice:", copySlice)
}