//package main
package common

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//按层级构造一棵二叉树
func CreateTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{nums[0],nil,nil}  // 根节点
	nodeQueue := []*TreeNode {root}
	for i:=1; i<len(nums); i=i+2 {  // 以 2 为单位移动
		if nodeQueue[0] == nil {   // 当前节点为 nil，其左、右子树均为 nil
			nodeQueue = append(nodeQueue, nil, nil)
			nodeQueue = nodeQueue[1:]  //
			continue
		}
		if nums[i] != 1000001 {   // 左子树不为 nil
			leftNode := &TreeNode{nums[i], nil,nil}
			nodeQueue[0].Left = leftNode
			nodeQueue = append(nodeQueue, leftNode)
		} else {  // 左子树为 nil
			nodeQueue = append(nodeQueue, nil)
		}

		if i+1 < len(nums) {   // 判断 nums 下标是否越界
			if nums[i+1] != 1000001 {   // 右子树不为 nil
				rightNode := &TreeNode{nums[i+1], nil,nil}
				nodeQueue[0].Right = rightNode
				nodeQueue = append(nodeQueue, rightNode)
			} else {   // 右子树为 nil
				nodeQueue = append(nodeQueue, nil)
			}
		}
		nodeQueue = nodeQueue[1:]
	}
	return root
}

//按层级遍历二叉树
func VisitNodeByLayer(root *TreeNode) {
	if root == nil {
		return
	}
	nodeQueue := []*TreeNode {root}   //存放所有节点的队列
	for len(nodeQueue) != 0 {
		if nodeQueue[0] != nil {
			fmt.Print(nodeQueue[0].Val, " ")
			nodeQueue = append(nodeQueue, nodeQueue[0].Left)
			nodeQueue = append(nodeQueue, nodeQueue[0].Right)
		} else {
			fmt.Print(nil, " ")
		}
		nodeQueue = nodeQueue[1:]
	}
}

// 1000001 表示 nil
//func main() {
//	root := createTree([]int {10,5,-3,3,2,1000001,11,3,-2,1000001,1})
//	visitNodeByLayer(root)
//}
