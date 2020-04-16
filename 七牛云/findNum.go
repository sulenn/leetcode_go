package main

import "fmt"

func findNum(nums []int, target int) int {
	if len(nums) == 0 || target<nums[0] || target > nums[len(nums)-1] {return -1}
	h := 0
	t:= len(nums)-1
	for h<=t {
		mid := h + (t-h)/2
		if nums[mid] == target {return mid}
		if nums[mid] < target {
			h = mid+1
		} else {
			t = mid
		}
	}
	return -1
}

func main() {
	fmt.Println(findNum([]int{1,2,3,4,5,6,7},4))
	fmt.Println(findNum([]int{1,2,3,4,5,6,7},-5))
	fmt.Println(findNum([]int{1,2,3,4,5,6,7},9))
	fmt.Println(findNum([]int{1,2,3,4,5,6,7},3))
	fmt.Println(findNum([]int{1,2,3,4,5,6,7},7))
	fmt.Println(findNum([]int{1,2,3,4,5,6,7},5))
}
