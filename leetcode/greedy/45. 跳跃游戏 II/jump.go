package main

//https://leetcode-cn.com/problems/jump-game-ii/

func jump(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {return 0}
	index := 0
	count := 1
	for index < len(nums) {
		num := nums[index]
		if num + index + 1 >= len(nums) {return count}
		max := index+1
		for i:=2; i<=num; i++ {
			if nums[max]+max-index < nums[index + i]+i {
				max = index + i
			}
		}
		index = max
		count++
	}
	return count
}
