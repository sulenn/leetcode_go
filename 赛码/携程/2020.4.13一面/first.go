package main

import "fmt"

func count(nums []int) int {
	//numMap := make(map[int]int)
	//	//for i:=0; i<len(nums); i++ {
	//	//	numMap[nums[i]]++
	//	//}
	//	//return numMap[7]

	arr := make([]int, 11)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]] += 1
	}
	return arr[7]
}

func main() {
	fmt.Println(count([]int{}))
	fmt.Println(count([]int{1, 2, 3, 4, 5}))
	fmt.Println(count([]int{1, 2, 3, 4, 5, 7, 7, 7, 7, 7}))
}
