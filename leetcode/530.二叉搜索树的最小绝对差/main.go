package main

import (
	"math"
)

//https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/

//参考：https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/solution/er-cha-sou-suo-shu-jie-dian-zui-xiao-ju-chi-by-lee/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min, prev int

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min = math.MaxInt64
	prev = math.MinInt64
	dfs(root)
	return min
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if prev != math.MinInt64 && root.Val-prev < min {
		min = root.Val - prev
	}
	prev = root.Val
	dfs(root.Right)
}
