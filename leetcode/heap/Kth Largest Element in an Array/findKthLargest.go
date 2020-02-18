package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums) - k]
}

////最大值沉底排序
//func findKthLargest(nums []int, k int) int {
//	if len(nums) == 0 || len(nums) < k {
//		return 0
//	}
//	for i:=0; i<k; i++ {
//		for j:=0; j<len(nums)-i-1;j++ {
//			if nums[j] > nums[j+1] {
//				nums[j], nums[j+1] = nums[j+1], nums[j]
//			}
//		}
//	}
//	return nums[len(nums) - k]
//}

func main() {
	fmt.Println(findKthLargest([]int {3,2,1,5,6,4}, 2))
	fmt.Println(findKthLargest([]int {3,2,3,1,2,4,5,5,6}, 4))
}
