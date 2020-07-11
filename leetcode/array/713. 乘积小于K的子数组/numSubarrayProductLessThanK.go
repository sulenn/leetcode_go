package main

//https://leetcode-cn.com/problems/subarray-product-less-than-k/

// 思路，用 mulSum 记录不超过 k 的最大连续子数组乘积，用 mulCount 记录该连续子数组中的元素数量
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k < 2 {
		return 0
	}
	mulSum := 1   // 不超过 k 的最大连续子数组乘积
	mulCount := 0 // 不超过 k 的最大连续子数组乘积中元素数量
	count := 0
	for i := 0; i < len(nums); i++ {
		mulSum *= nums[i]
		if mulSum < k { // 比较最大连续子数组乘积，若小于 k
			mulCount++
			count += mulCount
		} else { // 若大于 k
			for mulSum >= k { //累除，从连续子数组中的最小下标开始除
				if mulCount > 0 {
					mulSum /= nums[i-mulCount]
					mulCount--
				} else { // nums[i] 大于 k
					mulSum = 1
					mulCount = -1 // 注意这里是 -1
				}
			}
			mulCount++ // 注意加上 i 下标元素
			count += mulCount
		}
	}
	return count
}
