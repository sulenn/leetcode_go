package main

import (
	"container/list"
)

//https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	r := list.New()
	r.PushBack(root)
	result := make([]float64, 0)
	for r.Len() != 0 {
		temp := list.New()
		count := r.Len()
		sum := 0
		for r.Len() != 0 {
			ele := r.Front()
			curNode := ele.Value.(*TreeNode)
			r.Remove(ele)
			if curNode.Left != nil {
				temp.PushBack(curNode.Left)
			}
			if curNode.Right != nil {
				temp.PushBack(curNode.Right)
			}
			sum += curNode.Val
		}
		result = append(result, float64(sum)/float64(count))
		r = temp
	}
	return result
}
