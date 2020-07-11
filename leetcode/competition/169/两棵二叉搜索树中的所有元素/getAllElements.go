package main

//https://leetcode-cn.com/contest/weekly-contest-169/problems/all-elements-in-two-binary-search-trees/

//给你 root1 和 root2 这两棵二叉搜索树。
////
////请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.

//![image.png](https://ws1.sinaimg.cn/large/006alGmrgy1gakrbgumskj30z30qv400.jpg)

import (
	"../../../common"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *common.TreeNode, root2 *common.TreeNode) []int {
	result := []int{}
	var recursive func(root *common.TreeNode) []int
	recursive = func(root *common.TreeNode) []int { // 中序排序
		if root == nil {
			return []int{}
		}

		left := recursive(root.Left)
		res := append([]int{}, left...)
		res = append(res, root.Val)
		right := recursive(root.Right)
		res = append(res, right...)
		return res
	}
	root1List := recursive(root1)
	root2List := recursive(root2)
	root1Length := len(root1List) // 两棵树的有序长度
	root2Length := len(root2List)
	i, j := 0, 0
	for i < root1Length && j < root2Length { // 两个有序列表合成一个
		if root1List[i] < root2List[j] {
			result = append(result, root1List[i])
			i++
		} else {
			result = append(result, root2List[j])
			j++
		}
	}
	if i < root1Length {
		result = append(result, root1List[i:]...)
	} else {
		result = append(result, root2List[j:]...)
	}
	return result
}

func main() {
	fmt.Println(getAllElements(common.CreateTree([]int{2, 1, 4}), common.CreateTree([]int{1, 0, 3})))
	fmt.Println(getAllElements(common.CreateTree([]int{0, -10, 10}), common.CreateTree([]int{5, 1, 7, 0, 2})))
	fmt.Println(getAllElements(common.CreateTree([]int{0, -10, 10}), common.CreateTree([]int{})))
}
