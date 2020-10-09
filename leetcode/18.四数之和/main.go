package main

import "sort"

//https://leetcode-cn.com/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			a := j + 1
			b := len(nums) - 1
			for a < b {
				if nums[i]+nums[j]+nums[a]+nums[b] == target {
					result = append(result, []int{nums[i], nums[j], nums[a], nums[b]})
					for a < len(nums)-1 {
						a++
						if nums[a] != nums[a-1] {
							break
						}
					}
					for b > j {
						b--
						if nums[b] != nums[b+1] {
							break
						}
					}
				} else if nums[i]+nums[j]+nums[a]+nums[b] > target {
					for b > j {
						b--
						if nums[b] != nums[b+1] {
							break
						}
					}
				} else {
					for a < len(nums)-1 {
						a++
						if nums[a] != nums[a-1] {
							break
						}
					}
				}
			}
			for j < len(nums)-1 {
				j++
				if nums[j] != nums[j-1] {
					break
				}
			}
		}
		for i < len(nums)-1 {
			i++
			if nums[i] != nums[i-1] {
				break
			}
		}
	}
	return result
}
