package main

//https://leetcode-cn.com/problems/permutations/

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := helper(nums, []int{})
	return result
}

func helper(nums []int, tmp []int) [][]int {
	result := [][]int{}
	if len(nums) == 0 {
		result = append(result, tmp)
	}
	for v, k := range nums {
		newNums := []int{}
		if v != len(nums)-1 {
			newNums = append(newNums, nums[:v]...)
			newNums = append(newNums, nums[v+1:]...)
		} else {
			newNums = nums[:v]
		}
		result = append(result, helper(newNums, append(tmp, k))...)
	}
	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	//temp := []int {3}
	//fmt.Println(append(temp[:0], temp[0+1:]...))
	//competition := []int {1,2}
	//fmt.Println(competition, &competition[0])
	//temp := append(competition[:0],competition[1:]...)
	//fmt.Println(temp, &temp[0])
	//fmt.Println(competition, &competition[0])
}
