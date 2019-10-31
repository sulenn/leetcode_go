package main

import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return recursive(root, "")
}

func recursive(root *TreeNode, num string) int {
	if root.Left == nil && root.Right == nil {
		num += strconv.Itoa(root.Val)
		result,err:=strconv.Atoi(num)
		if err != nil {}
		return result
	}
	sum := 0
	if root.Left != nil {
		temp1 := num
		temp1 += strconv.Itoa(root.Val)
		sum += recursive(root.Left, temp1)
	}
	if root.Right != nil {
		temp2 := num
		temp2 += strconv.Itoa(root.Val)
		sum += recursive(root.Right, temp2)
	}
	return sum
}

func main() {
	
}
