package main

//https://leetcode-cn.com/problems/merge-two-binary-trees/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	leftTree := mergeTrees(t1.Left, t2.Left)
	rightTree := mergeTrees(t1.Right, t2.Right)
	t1.Val += t2.Val
	t1.Left = leftTree
	t1.Right = rightTree
	return t1
}

func main() {

}
