package main

//https://leetcode-cn.com/problems/partition-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var first *ListNode
	var second *ListNode
	var temp *ListNode
	temp.Next = head
	cur := head
	for cur != nil && cur.Val < x {
		first = cur
		cur = cur.Next
	}
}
