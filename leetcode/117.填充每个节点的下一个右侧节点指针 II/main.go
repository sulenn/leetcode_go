package main

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/

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
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			root.Left.Next = getRightNode(root)
		}
	}
	if root.Right != nil {
		root.Right.Next = getRightNode(root)
	}
	connect(root.Right) // 注意从右子树开始，因为上层的节点需要构建
	connect(root.Left)
	return root
}

func getRightNode(root *Node) *Node {
	temp := root.Next
	for temp != nil {
		if temp.Left != nil {
			return temp.Left
		} else if temp.Right != nil {
			return temp.Right
		}
		temp = temp.Next
	}
	return nil
}
