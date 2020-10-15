package main

//https://leetcode-cn.com/problems/delete-middle-node-lcci/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
