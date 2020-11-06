package main

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{Val: 0, Next: head}
	slow := pre
	fast := pre
	for n+1 != 0 && fast != nil {
		fast = fast.Next
		n--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}
