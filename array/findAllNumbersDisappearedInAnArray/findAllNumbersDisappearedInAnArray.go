package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++{
		index := 0
		if nums[i] < 0{
			index = -nums[i] -1
		} else {
			index = nums[i] -1
		}
		if nums[index] > 0 {
			nums[index] = - nums[index]
		}
	}
	var result []int
	for j:=0; j<len(nums); j++{
		if nums[j] > 0{
			result = append(result, j + 1)
		}
	}
	return result
}

func main() {
	in := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDisappearedNumbers(in))
}
