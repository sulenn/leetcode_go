package main

import "fmt"

//https://leetcode-cn.com/problems/diameter-of-binary-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var result int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result = 0
	recursive(root)
	return result
}
func recursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := recursive(root.Left)
	right := recursive(root.Right)
	if left + right > result {
		result = left + right
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func main() {
	var temp *TreeNode = &TreeNode{1,nil,nil}
	temp.Left = &TreeNode{2,&TreeNode{4,nil,nil}, &TreeNode{5,nil,nil}}
	temp.Right = &TreeNode{3,nil,nil}
	temp1 := &TreeNode{1,nil,nil}
	fmt.Println(diameterOfBinaryTree(temp))
	fmt.Println(diameterOfBinaryTree(temp1))
}
