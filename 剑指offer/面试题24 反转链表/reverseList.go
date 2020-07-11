package main

//https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		curNode := head
		head, curNode.Next = head.Next, prev
		prev = curNode
	}
	return prev
}
