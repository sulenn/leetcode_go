package main

import "fmt"

//https://leetcode-cn.com/problems/combinations/

var result [][]int

func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	backtracking(nums, k, []int{})
	return result
}

func backtracking(nums []int, k int, curArr []int) {
	if len(curArr) == k {
		result = append(result, curArr)
		return
	}
	if len(nums)+len(curArr) < k { // 剪枝
		return
	}
	for i := 0; i < len(nums); i++ {
		temp := make([]int, len(curArr))
		copy(temp, curArr)
		temp = append(temp, nums[i])
		backtracking(nums[i+1:], k, temp)
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
