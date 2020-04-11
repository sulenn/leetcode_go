package main

//https://leetcode-cn.com/problems/continuous-subarray-sum/

//Hash表解法参考：https://leetcode-cn.com/problems/continuous-subarray-sum/solution/lian-xu-de-zi-shu-zu-qiu-he-by-lenn123/
func checkSubarraySum(nums []int, k int) bool {
	cache := 0
	dic := make(map[int]struct{})
	sum := 0
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		if _,ok := dic[sum]; ok {
			return true
		}
		dic[cache] = struct{}{}
		cache = sum
	}
	return false
}

//时间复杂度O(n^2)
//func checkSubarraySum(nums []int, k int) bool {
//	if k == 0 {   // 特殊情况，（0，0）目标值为0
//		for i:=0; i<len(nums)-1; i++ {
//			if nums[i] == 0 && nums[i+1] == 0 {return true}
//		}
//		return false
//	}
//	if len(nums) == 0 {return false}
//	sum := make([]int, len(nums)+1)
//	for i:=0; i<len(nums); i++ {
//		sum[i+1] = sum[i] + nums[i]
//	}
//	for i:=0; i<len(nums); i++ {
//		for j:=i+2; j<len(nums)+1; j++ {    // 注意这儿 j = i+2，题目要求最少两个连续字符和
//			if (sum[j]-sum[i])%k == 0 {return true}
//		}
//	}
//	return false
//}
