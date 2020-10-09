package main

//https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后顺遍历
func convertBST(root *TreeNode) *TreeNode {
	createrTree(root, 0)
	return root
}

func createrTree(root *TreeNode, value int) int {
	if root == nil {
		return value
	}
	root.Val += createrTree(root.Right, value)
	return createrTree(root.Left, root.Val)
}
