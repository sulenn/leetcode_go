package main

import "sort"

//https://leetcode-cn.com/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	result := [][]int {}
	length := len(nums)
	sort.Ints(nums)
	if length <= 3 {
		return [][]int {}
	}
	for v:=0; v <length-3; v++ {  // 预留三位
		if v >= 1 && nums[v] == nums[v-1] {   //剪枝，去除重复的元素
			continue
		}
		//双指针，预留两位
		for i := v+1; i < length-2; i++ {
			pre := i + 1
			back := length - 1
			for pre < back {
				if nums[v] + nums[i] + nums[pre] + nums[back] == target {
					result = append(result, []int{nums[v], nums[i], nums[pre], nums[back]})
					for pre < back && nums[pre] == nums[pre+1] {  // 去除左边指针的重复元素
						pre++
					}
					for pre < back && nums[back] == nums[back-1] {  // 去除右边指针的重复元素
						back--
					}
					pre++
					back--
				} else if nums[v] + nums[i] + nums[pre] + nums[back] < target {
					pre ++
				} else {
					back --
				}
			}
			for i < length-2 && nums[i+1] == nums[i] {   //注意这儿是 nums[i+1] == nums[i] 剪枝，去除重复的元素
				i++
			}
		}
	}
	return result
}