package main

//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {return [][]int {}}
	curLevelNode := []*TreeNode {root}
	var curLevelValueArr []int
	var nextLevelNode []*TreeNode
	var result [][]int
	for len(curLevelNode) != 0 {
		curLevelValueArr = []int {}
		nextLevelNode = []*TreeNode {}
		for i:=0;i<len(curLevelNode);i++ {
			curLevelValueArr = append(curLevelValueArr, curLevelNode[i].Val)
			if curLevelNode[i].Left != nil {
				nextLevelNode = append(nextLevelNode, curLevelNode[i].Left)
			}
			if curLevelNode[i].Right != nil {
				nextLevelNode = append(nextLevelNode, curLevelNode[i].Right)
			}
		}
		if len(result) % 2 == 0{  // 当前层级为奇数
			result = append(result, curLevelValueArr)
		} else {  // 当前层级为偶数
			result = append(result, reverseArr(curLevelValueArr))
		}
		curLevelNode = nextLevelNode
	}
	return result
}

func reverseArr(nums []int) []int {   // 反转数组
	result := make([]int, len(nums))
	count := 0
	for i:=len(nums)-1; i>=0; i-- {
		result[count] = nums[i]
		count++
	}
	return result
}