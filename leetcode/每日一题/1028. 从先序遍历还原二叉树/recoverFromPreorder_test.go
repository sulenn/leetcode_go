package main

import (
	"fmt"
	"testing"
)

func Test_recoverFromPreorder(T *testing.T) {
	root := recoverFromPreorder("1-2--3--4-5--6--7")
	BinaryTreePrint(root)
	root = recoverFromPreorder("1-2--3---4-5--6---7")
	BinaryTreePrint(root)
	root = recoverFromPreorder("1-401--349---90--88")
	BinaryTreePrint(root)
}

func BinaryTreePrint(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	BinaryTreePrint(root.Left)
	BinaryTreePrint(root.Right)
}
