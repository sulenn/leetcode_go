package main

import "sort"

//核心思想，二分
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int,0)
	for i:=0; i<len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {   // 去重
			continue
		}
		head := i+1
		tail := len(nums)-1
		for head < tail {
			if head>i+1 && nums[head] == nums[head-1] {    // 去重
				head++
				continue
			}
			if nums[i]+nums[head]+nums[tail] == 0 {
				result = append(result, []int {nums[i], nums[head], nums[tail]})
				head++
				tail--
			} else if nums[i]+nums[head]+nums[tail] > 0 {
				tail--
			} else {
				head++
			}
		}
	}
	return result
}