package main

//https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var preOrder func(root *TreeNode) int
	preOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := preOrder(root.Left)
		right := preOrder(root.Right)
		if left == -1 || right == -1 || left-right > 1 || left-right < -1 {
			return -1
		}
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
	return preOrder(root) != -1
}
