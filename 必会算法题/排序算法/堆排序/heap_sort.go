package main

import "fmt"

func sort(nums []int) {
	length := len(nums)
	mid := length / 2
	for i := mid; i >= 0; i-- {
		for {
			left := i*2 + 1
			right := i*2 + 2
			if right < length && nums[left] < nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
			}
			if left < length && nums[i] < nums[left] {
				nums[i], nums[left] = nums[left], nums[i]
			} else {
				break
			}
			i = left
		}
	}
}

func main() {
	nums := []int{2, 1, 5, 6, 3, 4, 7, 2, 5}
	sort(nums)
	fmt.Println(nums)
}
