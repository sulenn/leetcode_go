package main

//https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/submissions/

import (
	"../../common"
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *common.TreeNode, k int) int {
	var smallestElement func(root *common.TreeNode)
	result := 0
	smallestElement = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		if k == 0 {   // 找到最小值即可，防止递归遍历所有节点
			return
		}
		smallestElement(root.Left)
		k--
		if k == 0{
			result = root.Val
		}
		smallestElement(root.Right)
	}
	smallestElement(root)
	return result
}

//
//func smallestElement(root *TreeNode) int {
//	if root
//
//}

func main() {
	root := common.CreateTree([]int {3,1,4,1000001,2})
	fmt.Println(kthSmallest(root,1))
	fmt.Println(kthSmallest(root,2))
	fmt.Println(kthSmallest(root,4))
}
