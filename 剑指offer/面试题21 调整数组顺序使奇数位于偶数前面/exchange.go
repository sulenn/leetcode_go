package main

//https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

func exchange(nums []int) []int {
	head := 0
	tail := len(nums) - 1
	for head < tail {
		if nums[head]&1 == 0 && nums[tail]&1 == 1 { // 头指针指向的元素为偶数；尾指针指向的元素为奇数
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail--
		}
		if nums[head]&1 == 1 { // 头指针指向的元素为奇数
			head++
		}
		if nums[tail]&1 == 0 { // 尾指针指向的元素为偶数
			tail--
		}
	}
	return nums
}
