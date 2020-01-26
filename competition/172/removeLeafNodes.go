package main

import (
	"../../common"
)

 //type TreeNode struct {
 //    Val int
 //    Left *TreeNode
 //    Right *TreeNode
 //}

 //递归
func removeLeafNodes(root *common.TreeNode, target int) *common.TreeNode {
	if root == nil {

	}
	var recursivr func(root *common.TreeNode, target int) bool
	recursivr = func(root *common.TreeNode, target int) bool {
		if root == nil {
			return true
		}
		left := recursivr(root.Left, target)
		if left {
			root.Left = nil
		}
		right := recursivr(root.Right, target)
		if right {
			root.Right = nil
		}
		if left && right && root.Val == target {
			return true
		} else {
			return false
		}
	}
	if recursivr(root,target) {   // 注意如果根结点满足条件，则根节点也需要删除
		root = nil
	}
	return root
}


func main() {
	root := common.CreateTree([]int {1,2,3,2,1000001,2,4})
	//removeLeafNodes(root, 2)
	common.VisitNodeByLayer(removeLeafNodes(root,2))
}
