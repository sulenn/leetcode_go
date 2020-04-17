package main

//https://leetcode-cn.com/problems/jump-game/

//该方法最简单（从前面开始遍历）：https://leetcode-cn.com/problems/jump-game/solution/55-by-ikaruga/

//思路，倒序遍历。如果遇到0，记录下标为zero，继续倒序遍历时下标需要满足nums[i] >= zere-i+1，才可以跳过这个zero所处的位置
//否则无法跳过这个0，即无法到达最后一个元素
func canJump(nums []int) bool {
	if len(nums) == 0 {return false}
	zere := -1
	for i:=len(nums)-2; i>=0; i-- {  // 从倒数第二个数开始
		if zere == -1 && nums[i] == 0 {
			zere = i
		}
		if zere != -1 && nums[i] >= zere-i+1 {
			zere = -1
		}
	}
	return zere==-1
}

//深搜超时
//func canJump(nums []int) bool {
//	if len(nums) == 0 {return false}
//	var dfs func(nums []int) bool
//	dfs = func(nums []int) bool {
//		if len(nums) == 1 || len(nums) == 0 {return true}
//		for i:=nums[0]; i>=1; i-- {
//			if 1+i <= len(nums) {
//				if dfs(nums[i:]) {return true}
//			} else {return true}
//		}
//		return false
//	}
//	return dfs(nums)
//}

