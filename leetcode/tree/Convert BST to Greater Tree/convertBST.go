package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	recursive(root)
	return root
}

func recursive(root *TreeNode) {
	if root != nil {
		recursive(root.Right)
		sum += root.Val
		root.Val = sum
		recursive(root.Left)
	}
}

func main() {
	temp := &TreeNode{5, &TreeNode{2, nil, nil}, &TreeNode{13, nil, nil}}
	fmt.Println(convertBST(temp))
}
