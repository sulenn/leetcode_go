package main

//https://leetcode-cn.com/problems/path-sum-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	dfs(root, sum, []int{}, &result)
	return result
}

func dfs(root *TreeNode, target int, nums []int, result *[][]int) {
	if root == nil {
		return
	}
	target -= root.Val
	nums = append(nums, root.Val)
	if root.Left == nil && root.Right == nil && target == 0 {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
	}
	dfs(root.Left, target, nums, result)
	dfs(root.Right, target, nums, result)
}
