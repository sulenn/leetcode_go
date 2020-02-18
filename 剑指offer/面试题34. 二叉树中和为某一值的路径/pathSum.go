package main

import "fmt"

//https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	//if root == nil || root.Val > sum {return [][]int {}}
	if root == nil {return [][]int {}}
	if root.Left == nil && root.Right == nil && sum == root.Val {  // 满足条件的路径
		return [][]int {{root.Val}}
	}
	if root.Left == nil && root.Right == nil && sum != root.Val {
		return [][]int {}
	}
	leftTreePath := [][]int {}  // 左树路径
	rightTreePath := [][]int {}  // 右树路径
	if root.Left != nil {
		leftTreePath = pathSum(root.Left, sum-root.Val)
	}
	for i:=0;i<len(leftTreePath);i++ {
		leftTreePath[i] = append([]int {root.Val}, leftTreePath[i]...)
	}
	if root.Right != nil {
		rightTreePath = pathSum(root.Right, sum-root.Val)
	}
	for i:=0;i<len(rightTreePath);i++ {
		rightTreePath[i] = append([]int {root.Val}, rightTreePath[i]...)
	}
	return append(leftTreePath, rightTreePath...)
}
