package main

import "strconv"

//https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/

func evalRPN(tokens []string) int {
	nums := make([]int, 0)
	for i:=0; i<len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			if len(nums) < 2 {return 0}
			if tokens[i] == "+" {
				nums[len(nums)-2] += nums[len(nums)-1]
			}
			if tokens[i] == "-" {
				nums[len(nums)-2] -= nums[len(nums)-1]
			}
			if tokens[i] == "*" {
				nums[len(nums)-2] *= nums[len(nums)-1]
			}
			if tokens[i] == "/" {
				nums[len(nums)-2] /= nums[len(nums)-1]
			}
			nums = nums[:len(nums)-1]
		} else {
			num, err := strconv.Atoi(tokens[i])
			if err != nil {
				return 0
			}
			nums = append(nums, num)
		}
	}
	return nums[len(nums)-1]
}