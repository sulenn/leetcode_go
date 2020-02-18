package main

//https://leetcode-cn.com/problems/distribute-coins-in-binary-tree/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//递归后序深搜
//参考：https://leetcode-cn.com/problems/distribute-coins-in-binary-tree/solution/zai-er-cha-shu-zhong-fen-pei-ying-bi-by-leetcode/
func distributeCoins(root *TreeNode) int {
	result := 0
	var recursive func(root *TreeNode) int
	recursive = func(root *TreeNode) int {
		if root == nil {return 0}
		left := recursive(root.Left)
		right := recursive(root.Right)
		result += abs(left) + abs(right)
		return root.Val + left + right - 1
	}
	recursive(root)
	return result
}

func abs(n int) int {
	if n < 0 {return -n}
	return n
}

func main() {
}
