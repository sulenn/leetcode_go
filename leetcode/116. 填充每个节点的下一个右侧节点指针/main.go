package main

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil && root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
