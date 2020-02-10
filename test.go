package main

import "fmt"

func main() {
	fmt.Println(1<<1)
	fmt.Println(1<<2)
	fmt.Println(1<<3)
	var temp [][]int
	//temp[0] = []int {1}   // 报错，没有初始化，temp 默认长度尾 0
	fmt.Println(temp)
	fmt.Println(len(temp))
}