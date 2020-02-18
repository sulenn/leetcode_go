package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := 0
	for _, elem := range nums {
		if elem != val {
			nums[length] = elem
			length++
		}
	}
	return length
}

func main() {
	fmt.Println(removeElement([]int{3,2,2,3}, 3))
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
}
