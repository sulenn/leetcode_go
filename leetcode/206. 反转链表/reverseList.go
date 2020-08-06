package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curNode := head
	nextNode := head.Next
	curNode.Next = nil
	for nextNode != nil {
		temp := nextNode.Next
		nextNode.Next = curNode
		curNode, nextNode = nextNode, temp
	}
	return curNode
}
