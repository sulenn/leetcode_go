package main

//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/

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
		result = append(result, curLevelValueArr)
		curLevelNode = nextLevelNode
	}
	return result
}
