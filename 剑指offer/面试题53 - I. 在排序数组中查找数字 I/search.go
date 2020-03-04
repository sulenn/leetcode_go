package main

//https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

// 二分法找一个和最后一个小标
func search(nums []int, target int) int {
	first := firstK(nums, target)
	last := lastK(nums, target)
	if first == -1 || last == -1 {
		return 0
	}
	return last-first+1
}

func firstK(nums []int, target int) int {   // 二分法找第一个 target 下标
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				end = mid - 1
			}
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func lastK(nums []int, target int) int {   // 二分法找最后一个 target 下标
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				start = mid + 1
			}
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
