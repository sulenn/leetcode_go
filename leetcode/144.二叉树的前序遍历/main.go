package main

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//	result := make([]int, 0)
//	result = append(result, root.Val)
//	left := preorderTraversal(root.Left)
//	result = append(result, left...)
//	right := preorderTraversal(root.Right)
//	result = append(result, right...)
//	return result
//}

// 迭代
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nodeList := []*TreeNode{root}
	result := make([]int, 0)
	for len(nodeList) != 0 {
		node := nodeList[len(nodeList)-1]
		nodeList = nodeList[:len(nodeList)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			nodeList = append(nodeList, node.Right)
		}
		if node.Left != nil {
			nodeList = append(nodeList, node.Left)
		}
	}
	return result
}
