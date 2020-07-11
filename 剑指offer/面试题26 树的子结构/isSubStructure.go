package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val {
		if doesBTreeHaveinATree(A, B) {
			return true
		}
	}
	leftTree := isSubStructure(A.Left, B)
	rightTree := isSubStructure(A.Right, B)
	return leftTree || rightTree
}

func doesBTreeHaveinATree(A *TreeNode, B *TreeNode) bool { // 判断 A 中是否存在 B 子树
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return doesBTreeHaveinATree(A.Left, B.Left) && doesBTreeHaveinATree(A.Right, B.Right)
}
