package main

//https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, value int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return value*10 + root.Val
	}
	return dfs(root.Left, value*10+root.Val) + dfs(root.Right, value*10+root.Val)
}
