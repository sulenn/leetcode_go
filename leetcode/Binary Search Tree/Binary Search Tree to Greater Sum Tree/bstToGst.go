package main

//https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 从右子树开始，累加
func bstToGst(root *TreeNode) *TreeNode {
	var recursive func(root *TreeNode)
	sum := 0
	recursive = func(root *TreeNode) {
		if root == nil {
			return
		}
		recursive(root.Right)
		sum += root.Val
		root.Val = sum
		recursive(root.Left)
	}
	recursive(root)
	return root
}

func main() {
	
}
