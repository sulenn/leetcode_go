package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Val   int
	lines int
}

//思路：利用 Items 记录每一个节点所在的层，然后对 Items 进行 DFS 深搜
func recoverFromPreorder(S string) *TreeNode {
	items := make([]Item, 0)
	lineNums := 0
	numStr := ""
	for i := 0; i < len(S); i++ {
		if S[i] == '-' && numStr != "" {
			num, _ := strconv.Atoi(numStr)
			items = append(items, Item{num, lineNums})
			lineNums = 1
			numStr = ""
		} else if S[i] == '-' {
			lineNums++
		} else {
			numStr += S[i : i+1]
		}
	}
	num, _ := strconv.Atoi(numStr)
	items = append(items, Item{num, lineNums})
	nums, root := dfs(0, 0, items)
	fmt.Println(nums)
	return root
}

// nums 用于记录遍历的节点数，layer 表示当前遍历的树层级
func dfs(layer int, nums int, items []Item) (int, *TreeNode) {
	if len(items) == nums || items[nums].lines != layer {
		return nums, nil
	}
	node := &TreeNode{items[nums].Val, nil, nil}
	nums++
	nums, node.Left = dfs(layer+1, nums, items)
	nums, node.Right = dfs(layer+1, nums, items)
	return nums, node
}
