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

func quickSort(nums []int) { // 真正的快排
	if len(nums) == 0 {
		return
	}
	l := 0
	r := len(nums) - 1
	partition(nums, l, r)
}

func partition(nums []int, l, r int) { // 快排精髓
	if l >= r {
		return
	}
	v := nums[l]
	less := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			less++
			nums[less], nums[i] = nums[i], nums[less]
		}
	}
	nums[l], nums[less] = nums[less], nums[l]
	partition(nums, l, less-1)
	partition(nums, less+1, r)
}
