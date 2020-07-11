package main

//https://leetcode-cn.com/problems/house-robber-ii/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max1 := rob1(nums[:len(nums)-1]) // 不包括最后的房子
	max2 := rob1(nums[1:])           // 不包括第一个房子
	if max1 > max2 {
		return max1
	}
	return max2
}

func rob1(nums []int) int {
	pre1, pre2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if pre1+nums[i] > pre2 {
			pre2, pre1 = pre1+nums[i], pre2
		} else {
			pre1 = pre2
		}
	}
	return pre2
}
