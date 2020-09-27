package main

import "fmt"

//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node, _ := dfs(root, p, q)
	return node
}

func dfs(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	leftNode, leftFlag := dfs(root.Left, p, q)
	if leftNode != nil {
		return leftNode, true
	}
	rightNode, rightFlag := dfs(root.Right, p, q)
	if rightNode != nil {
		return rightNode, true
	}
	if leftFlag && rightFlag {
		return root, true
	}
	if (root == p || root == q) && (leftFlag || rightFlag) { // 当前节点即是最近公共父节点
		return root, true
	}
	if root == p || root == q {
		return nil, true
	}
	return nil, leftFlag || rightFlag
}

func main() {
	node1 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 8}
	node4 := &TreeNode{Val: 0}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 7}
	node7 := &TreeNode{Val: 9}
	node8 := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 5}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Left = node8
	node5.Right = node9
	node := lowestCommonAncestor(node1, node2, node5)
	fmt.Println(node)
}
