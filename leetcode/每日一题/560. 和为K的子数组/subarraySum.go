package main

//https://leetcode-cn.com/problems/subarray-sum-equals-k/

//前缀和，利用 hashmap 保存计算的中间结果
//参考：https://leetcode-cn.com/problems/subarray-sum-equals-k/solution/qian-zhui-he-si-xiang-560-he-wei-kde-zi-shu-zu-by-/
func subarraySum(nums []int, k int) int {
	count := 0
	sumsMap := make(map[int]int)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == k {
			count++
		}
		if value, ok := sumsMap[sum-k]; ok {
			count += value
		}
		sumsMap[sum]++
	}
	return count
}
