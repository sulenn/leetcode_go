package main

//https://leetcode-cn.com/problems/binary-tree-cameras/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//参考：https://leetcode-cn.com/problems/binary-tree-cameras/solution/chao-ji-hao-li-jie-de-da-an-by-levyjeng/
func minCameraCover(root *TreeNode) int {
	count, flag := dfs(root)
	if flag == 0 {
		count++
	}
	return count
}

// 0 表示父亲或者子节点中没有摄像头
// 1 表示父亲或者子节点中有摄像头
// 2 表示当前节点中有摄像
func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 1
	}
	count1, flag1 := dfs(root.Left)
	count2, flag2 := dfs(root.Right)
	if flag1 == 0 || flag2 == 0 {
		return count1 + count2 + 1, 2
	}
	if flag1 == 1 && flag2 == 1 {
		return count1 + count2, 0
	}
	return count1 + count2, 1
}
