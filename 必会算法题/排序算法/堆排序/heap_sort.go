package main

import "fmt"

func heapSort(nums []int) {
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		adjustHeap(nums, i)
	}
}

func adjustHeap(nums []int, pos int) {
	length := len(nums)
	for {
		child := pos*2 + 1
		if child+1 < length && nums[child+1] > nums[child] {
			child++
		}
		if child < length && nums[child] > nums[pos] {
			nums[child], nums[pos] = nums[pos], nums[child]
		} else {
			break
		}
		pos = child
	}
}

func main() {
	//nums := []int{2, 1, 5, 6, 3, 4, 7, 2, 5}
	nums := []int{2, 1, 4, 5, 3, 7, 8, 6}
	heapSort(nums)
	fmt.Println(nums)
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums[:i], 0)
	}
	fmt.Println(nums)
}
