package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 0
	for _  , elem := range nums {
		if nums[length] != elem {
			nums[length + 1] = elem
			length++
		}
	}
	return length + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1,1,2}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}
