package main

//https://leetcode-cn.com/problems/the-masseuse-lcci/

//该题和 “198. 打家劫舍” 一模一样
func massage(nums []int) int {
	first, second := 0, 0
	for _, v := range nums {
		if first+v > second {
			second, first = first+v, second
		} else {
			first = second
		}
	}
	return second
}
