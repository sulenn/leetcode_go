package main

import "fmt"

//https://leetcode-cn.com/problems/duplicate-zeros/

//两次遍历，第一次重头开始判断 0 的数量。第二次从后往前遍历、覆盖
func duplicateZeros(arr []int) {
	zeroNums := 0
	i := 0        // 第一次遍历的最后一个下标
	l := len(arr) // 数组长
	for ; i < len(arr)-zeroNums; i++ {
		if arr[i] == 0 {
			zeroNums++
		}
	}
	if zeroNums+i-1 == l { // 第一次遍历时最后一个元素为 0
		arr[l-1] = arr[i-1]
		i--
		l--
	}
	i--
	l--
	for ; i >= 0; i-- {
		arr[l] = arr[i]
		l--
		if arr[i] == 0 {
			arr[l] = arr[i]
			l--
		}
	}
	fmt.Println(arr)
}
