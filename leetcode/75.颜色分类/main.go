package main

import "fmt"

//https://leetcode-cn.com/problems/sort-colors/

func sortColors(nums []int) {
	p1, p2, p3 := 0, 0, len(nums)-1
	for p2 <= p3 {
		if nums[p2] == 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1++
			p2++ // 注意当前指针 p2++
		} else if nums[p2] == 2 {
			nums[p2], nums[p3] = nums[p3], nums[p2]
			p3--
		} else {
			p2++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
