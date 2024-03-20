package main

import (
	"fmt"
)

func test() {
	var arr [3]int
	brr := arr[1:2]      // 创建一个切片，引用数组 arr。长度为 1，容量为 2
	brr = append(brr, 9) // 长度 为 2， 容量为 2
	brr = append(brr, 9) // 长度为 3 ，容量为 4
	fmt.Println(len(brr), cap(brr))
}

func main() {
	test()
}
