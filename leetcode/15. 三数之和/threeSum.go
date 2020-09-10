package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/3sum/

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	p1 := 0
	p2 := 0
	p3 := 0
	length := len(nums)
	for p1+2 < length {
		p2 = p1 + 1
		p3 = length - 1
		for p2 < p3 {
			if nums[p1]+nums[p2]+nums[p3] == 0 {
				result = append(result, []int{nums[p1], nums[p2], nums[p3]})
			}
			if nums[p1]+nums[p2]+nums[p3] > 0 {
				p3--
				continue
			}
			if nums[p1]+nums[p2]+nums[p3] < 0 {
				p2++
				continue
			}
			p2++
			for p2 < length && nums[p2] == nums[p2-1] { // 去重
				p2++
			}
			p3--
			for p3 > 1 && nums[p3] == nums[p3+1] { //
				p3--
			}
		}
		p1++
		for p1+2 < length && nums[p1] == nums[p1-1] {
			p1++
		}
	}
	return result
}

func main() {
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{1, -1, -1, 0}))
}
