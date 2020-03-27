package main

//https://leetcode-cn.com/problems/non-decreasing-array/

//考虑3、2、2和3、3、2两种情况
func checkPossibility(nums []int) bool {
	if len(nums) < 2 {return true}
	num := 1
	if nums[0]>nums[1] {
		nums[0] = nums[1]
		num--
	}
	for i:=1 ;i<len(nums) -1 && num >= 0 ; i++ {
		if nums[i]>nums[i+1] && nums[i-1] > nums[i+1] {  // 3，3，2
			nums[i+1]=nums[i]
			num--
			continue
		}
		if nums[i]>nums[i+1] {  // 3，2，2
			nums[i]=nums[i+1]
			num--
			continue
		}
	}
	return num>=0
}
