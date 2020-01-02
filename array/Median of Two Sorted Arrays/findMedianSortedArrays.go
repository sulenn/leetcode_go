package main

import "fmt"

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newNums := []int {}   // 新排序切片
	len1 := 0
	len2 := 0
	for len1 < len(nums1) && len2 < len(nums2) {  // 两个已排序切片变为一个排序切片
		if nums1[len1] < nums2[len2] {
			newNums = append(newNums, nums1[len1])
			len1++
		} else {
			newNums = append(newNums, nums2[len2])
			len2++
		}
	}
	if len1 < len(nums1) {    // 剩下的切片
		newNums = append(newNums, nums1[len1:]...)
	} else {
		newNums = append(newNums, nums2[len2:]...)
	}
	mid := len(newNums) / 2
	if len(newNums) % 2 == 0 {  // 单双数，影响中位数取值
		return float64((newNums[mid - 1] + newNums[mid])) / 2.0
	} else {
		return float64(newNums[mid])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int {1, 3}, []int {2}))
	fmt.Println(findMedianSortedArrays([]int {1, 2}, []int {3, 4}))
}
