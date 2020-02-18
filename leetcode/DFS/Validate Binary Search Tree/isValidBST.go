package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return recursive(root, math.MaxInt64, math.MinInt64)
}

func recursive(root *TreeNode, upper int, low int) bool {
	if root == nil {
		return true
	}
	if root.Val >= upper || root.Val <= low {
		return false
	}
	return recursive(root.Left, root.Val, low) && recursive(root.Right, upper, root.Val)
}

func main() {
	
}
