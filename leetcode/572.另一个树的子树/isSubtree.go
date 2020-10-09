package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	var dfs func(root1 *TreeNode, root2 *TreeNode) bool
	dfs = func(root1 *TreeNode, root2 *TreeNode) bool {
		if root1 == nil && root2 == nil {
			return true
		}
		if root1 == nil || root2 == nil {
			return false
		}
		if root1.Val == root2.Val { // 值相等
			if equalTree(root1, root2) {
				return true
			} // 判断两颗树是否相等
		}
		// 如果 root1.val 不等于 root2.val，则递归 root1 的左右子树
		// 如果 root1.val 等于 root2.val，但是两颗树不相等，则递归 root1 的左右子树
		return dfs(root1.Left, t) || dfs(root1.Right, t)
	}
	return dfs(s, t)
}

func equalTree(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return equalTree(root1.Left, root2.Left) && equalTree(root1.Right, root2.Right)
}
