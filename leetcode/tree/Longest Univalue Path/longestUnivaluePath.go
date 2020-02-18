package main

//https://leetcode-cn.com/problems/longest-univalue-path/

//我们可以将任何路径（具有相同值的节点）看作是最多两个从其根延伸出的箭头。
//
//具体地说，路径的根将是唯一节点，因此该节点的父节点不会出现在该路径中，而箭头将是根在该路径中只有一个子节点的路径。
//
//然后，对于每个节点，我们想知道向左延伸的最长箭头和向右延伸的最长箭头是什么？我们可以用递归来解决这个问题。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var recursive func(root *TreeNode) int
	result := 0
	recursive = func(root *TreeNode) int {   // 每个节点的最长路径等于左边最新最长路径 + 右边最新最长路径
		if root == nil {
			return 0
		}
		left := recursive(root.Left)  // 左子树最长路径
		right := recursive(root.Right)   // 右子树最长路径
		arrowLeft, arrowRight := 0, 0   // 包括节点在内的左边、右边最长路径
		if root.Left != nil && root.Left.Val == root.Val {   // 计算左边
			arrowLeft = left + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {   // 计算右边
			arrowRight = right + 1
		}
		result = max(result, arrowLeft + arrowRight)
		return max(arrowLeft, arrowRight)
	}
	recursive(root)
	return result
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	
}
