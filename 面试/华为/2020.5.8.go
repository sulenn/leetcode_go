package main

import "fmt"

func maxUpArr(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	countNum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] {
				if countNum[i] < countNum[j]+1 {
					countNum[i] = countNum[j] + 1
				}
			}
		}
		if countNum[i] == 0 {
			countNum[i] = 1
		}
	}
	return countNum[len(nums)-1]
}

func main() {
	fmt.Println(maxUpArr([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(maxUpArr([]int{10, 9, 2, 5, 3, 4, 7, 101, 18}))
	fmt.Println(maxUpArr([]int{2, 2, 2, 2, 2, 2}))
	fmt.Println(maxUpArr([]int{2}))
	fmt.Println(maxUpArr([]int{}))
}
