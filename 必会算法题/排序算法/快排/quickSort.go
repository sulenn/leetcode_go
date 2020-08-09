package main

func sort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	fixed := 0
	head := 1
	tail := len(nums) - 1
	for head <= tail {
		if fixed < tail {
			if nums[fixed] > nums[tail] {
				nums[fixed], nums[tail] = nums[tail], nums[fixed]
				fixed = tail
			}
			tail--
		} else {
			if nums[fixed] < nums[head] {
				nums[fixed], nums[head] = nums[head], nums[fixed]
				fixed = head
			}
			head++
		}
	}
	sort(nums[:fixed])
	sort(nums[fixed+1:])
}
