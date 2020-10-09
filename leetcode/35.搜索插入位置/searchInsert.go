package main

//https://leetcode-cn.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1
	var mid int
	for head <= tail {
		mid = head + (tail-head)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}
	return head
}
