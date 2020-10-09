package main

//https://leetcode-cn.com/problems/middle-of-the-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	p := head
	for p.Next != nil && p.Next.Next != nil {
		head = head.Next
		p = p.Next.Next
	}
	if p.Next != nil {
		return head.Next
	}
	return head
}
