package main

//https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	if len(nums)==0 {return []int {-1,-1}}
	l1 := 0
	h1 := len(nums)-1
	for l1 < h1 {  // 二分查找最左边
		mid := l1 + (h1-l1)/2
		if nums[mid] < target {
			l1 = mid+1
		} else {
			h1 = mid
		}
	}
	if nums[h1] != target {return []int {-1,-1}}
	l2 := 0
	h2 := len(nums)-1
	for l2 < h2 {  // 二分查找最右边
		mid := l2 + (h2-l2)/2
		if nums[mid+1] <= target {
			l2 = mid+1
		} else {
			h2 = mid
		}
	}
	return []int{h1,h2}
}