package main

//https://leetcode-cn.com/problems/rotate-array/

//参考：https://leetcode-cn.com/problems/rotate-array/solution/javadai-ma-3chong-fang-shi-tu-wen-xiang-q8lz9/
//多次翻转
func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	reverse(nums, 0, length-1) // 全部翻转
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
