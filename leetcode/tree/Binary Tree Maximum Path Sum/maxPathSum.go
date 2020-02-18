package main

import (
	"../../common"
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//这道题和 Longest Univalue Path 的解题思路一样，每个节点都分左右 path。
//节点的最大 path 等于左边临近的最大 Path 加上 右边临近的最大 path 加上节点本身 value
//采用后序遍历的方法

func maxPathSum(root *common.TreeNode) int {
	var recursive func(root *common.TreeNode) int
	maxSum := math.MinInt64    // 最小值
	recursive = func(root *common.TreeNode) int {
		if root == nil {
			return 0
		}
		left := recursive(root.Left)
		right := recursive(root.Right)
		if left < 0 {    // 如果左边小于0, 变为 0， 不能让负数降低节点最大路径的值
			left = 0
		}
		if right < 0 {  // 如果右边小于 0
			right = 0
		}
		maxSum = max(maxSum, left + right + root.Val)
		return max(left + root.Val, right + root.Val)    // 节点（包括节点在内）左边和右边中较大的值
	}
	recursive(root)
	return maxSum
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(math.MaxInt64, math.MinInt64)
	root := common.CreateTree([]int {2, -1})
	fmt.Println(maxPathSum(root))
}
