package main

import (
	"../../common"
	"fmt"
)

//https://leetcode-cn.com/problems/path-sum-iii/

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func pathSum(root *common.TreeNode, sum int) int {
	var f func(root *common.TreeNode, sums []int)
	var ans int

	f = func(root *common.TreeNode, path []int) {
		if root == nil {
			return
		}

		path = append(path, root.Val)

		tmp := 0
		for i := len(path)-1; i >= 0; i-- {
			tmp += path[i]
			if tmp == sum {
				ans++
			}
		}

		f(root.Left, path)
		f(root.Right, path)
	}

	f(root, make([]int, 0, 1000))
	return ans
}

//第二种解法
//func pathSum(root *common.TreeNode, sum int) int {
//	return recursive(root, []int {}, sum)
//}
//
//func recursive(root *common.TreeNode,num []int,sum int) int{
//	if root==nil {
//		return 0
//	}
//	res :=0
//	num = append(num, root.Val)
//	temp := 0  // 累加值
//	for i:=len(num)-1;i>=0;i--  {     //从靠近叶子节点的节点开始累加
//		temp+=num[i]
//		if temp==sum{
//			res++
//		}
//	}
//	res += recursive(root.Left, num, sum)
//	res += recursive(root.Right, num, sum)
//	return res
//}

//第一种解法
//func recursive(root *common.TreeNode, path []int, sum int) int {
//	if root == nil {
//		return 0
//	}
//	result := 0
//	newPath := append([]int {}, path...)    // 递归调用时，对切片内容的修改会带来影响。如：左递归中切片值修改会保留至右递归中
//	for i:=0; i<len(newPath); i++ {
//		if newPath[i] + root.Val == sum{
//			result += 1
//		}
//		newPath[i] += root.Val
//	}
//	if root.Val == sum {
//		result ++
//	}
//	newPath = append(newPath, root.Val)
//	result += recursive(root.Left, newPath, sum)
//	result += recursive(root.Right, newPath, sum)
//	return result
//}

func main() {
	root := common.CreateTree([]int {10,5,-3,3,2,1000001,11,3,-2,1000001,1})
	//common.VisitNodeByLayer(root)
	fmt.Println(pathSum(root, 8))
}
