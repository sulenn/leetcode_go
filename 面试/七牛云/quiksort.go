package main

import "fmt"

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	p := 0
	h := 1
	t := len(nums) - 1
	for h <= t {
		if p < t {
			if nums[p] > nums[t] {
				nums[p], nums[t] = nums[t], nums[p]
				p = t
			}
			t--
		}
		if p > h {
			if nums[p] < nums[h] {
				nums[p], nums[h] = nums[h], nums[p]
				p = h
			}
			h++
		}
	}
	quickSort(nums[:p])
	quickSort(nums[p+1:])
}

func main() {
	test := []int{4, 2, 8, 1, 6, 9, 5, 7}
	quickSort(test)
	fmt.Println(test)
	test = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	quickSort(test)
	fmt.Println(test)
}
