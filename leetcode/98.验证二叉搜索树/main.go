package main

import "math"

//https://leetcode-cn.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MaxInt64, math.MinInt64)
}

func dfs(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return dfs(root.Left, root.Val, min) && dfs(root.Right, max, root.Val)
}
