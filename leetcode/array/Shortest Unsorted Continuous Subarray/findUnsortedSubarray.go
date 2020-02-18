package main

import "fmt"

//https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/

func findUnsortedSubarray(nums []int) int {
	head := 0
	tail := len(nums) - 1
	flag := false
	for head < len(nums) {
		for i := head+1; i < len(nums); i++ {
			if nums[i] < nums[head] {
				head--
				flag = true
				break
			}
		}
		if flag {
			break
		}
		head++
	}
	flag = false
	for tail >= 0 {
		for i:=tail - 1; i >= 0; i-- {
			if nums[i] > nums[tail] {
				tail++
				flag = true
				break
			}
		}
		if flag {
			break
		}
		tail--
	}
	if tail <= head {
		return 0
	}
	return tail - head -1
}

func main() {
	fmt.Println(findUnsortedSubarray([]int {2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int {2}))
	fmt.Println(findUnsortedSubarray([]int {}))
}
