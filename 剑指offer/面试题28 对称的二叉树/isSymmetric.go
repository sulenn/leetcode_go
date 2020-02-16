package main

//https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func isSymmetric(root *TreeNode) bool {  // 判断树的中序和后序是否相等即可
	return judgeSymmetric(root,root)
}

func judgeSymmetric(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {return true}
	if A == nil || B == nil {return false}
	if A.Val != B.Val {return false}
	return judgeSymmetric(A.Left, B.Right) && judgeSymmetric(A.Right, B.Left)
}
