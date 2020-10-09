package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	root = dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftTree := dfs(root.Left)
	rightTree := dfs(root.Right)
	if leftTree == nil {
		root.Right = rightTree
	} else {
		root.Right = leftTree
		for leftTree.Right != nil {
			leftTree = leftTree.Right
		}
		leftTree.Right = rightTree
		root.Left = nil
	}
	return root
}
