package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/

// 更简单的写法，不需要传大量指针：https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/xian-lai-ge-zhong-xu-di-gui-man-man-bu-zuo-ye-by-l/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	result := &[]int{}
	max, curValue := 0, 0
	dfs(root, &TreeNode{Val: math.MinInt64}, result, &max, &curValue)
	return *result
}

func dfs(root *TreeNode, pre *TreeNode, result *[]int, max *int, curValue *int) {
	if root == nil {
		return
	}
	dfs(root.Left, pre, result, max, curValue) // 左
	if pre.Val == root.Val {                   // 中
		*curValue++
	} else {
		*curValue = 1
	}
	if *curValue > *max {
		*result = make([]int, 0)
		*result = append(*result, root.Val)
		*max = *curValue
	} else if *curValue == *max {
		*result = append(*result, root.Val)
	}
	*pre = *root
	dfs(root.Right, pre, result, max, curValue) // 右
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	//node3 := &TreeNode{Val: 2}
	//node1.Right = node2
	//node2.Left = node3
	node1.Left = node2
	fmt.Println(findMode(node1))
}
