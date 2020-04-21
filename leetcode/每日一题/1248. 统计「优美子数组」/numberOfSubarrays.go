package main

//https://leetcode-cn.com/problems/count-number-of-nice-subarrays/

//双指针，思路：
//1. p1为0时，p1和p2指向满足条件的近端和远端数组区间，此时满足条件（k个奇数）的连续数组数量有n
//2. p1+1，判断nums[p1]，如果nums[p1]为偶数，那么p2不动，此时满足条件的连续数组数量任然是n
//	如果nums[p1]为奇数，此时p1+1到p2区间中存在k-1个奇数，累加p2，找到满足条件且最远端p2+x，此时p1+1到p2+x区间的数量为m
//3. 重复第2个过程。
func numberOfSubarrays(nums []int, k int) int {
	p1 := 0
	p2 := 0
	curLevelCount := 0
	count := 0
	num := 0
	for ;p1 <= len(nums) - k; p1++{
		if p1 != 0 && nums[p1-1]%2 == 0 {  // nums[p1] 为偶数
			count += curLevelCount
			continue
		} else if p1 != 0 && nums[p1-1]%2 != 0 {  // nums[p1] 为奇数
			curLevelCount = 0
			num = k-1
		}
		for ;p2 < len(nums); p2++ {
			if nums[p2] % 2 != 0 {num++}
			if num == k {curLevelCount++}
			if num > k {break}
		}
		count += curLevelCount
	}
	return count
}
