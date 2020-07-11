package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result *TreeNode

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	recursive(root, p, q)
//	return result
//}
//
//func recursive(root, p, q *TreeNode) bool {
//	if root == nil {
//		return false
//	}
//	left := recursive(root.Left, p, q)
//	right := recursive(root.Right, p, q)
//	mid := (root == p) || (root == q)
//	if (left && right) || (left && mid) || (mid && right) {
//		result = root
//	}
//	return mid || left || right
//}

//another way to solve this question, more easy and convenient
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}
