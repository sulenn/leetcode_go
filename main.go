package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 6, 7, 3, 5, 8, 3, 9}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int, start, end int) {
	if end >= len(nums) || start >= end {
		return
	}
	fixValue, point := start, start
	for point <= end {
		if nums[point] > nums[start] {
			fixValue++
			nums[fixValue], nums[point] = nums[point], nums[fixValue]
		}
		point++
	}
	nums[start], nums[fixValue] = nums[fixValue], nums[start]
	quickSort(nums, start, fixValue-1)
	quickSort(nums, fixValue+1, end)
}
