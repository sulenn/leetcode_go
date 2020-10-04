package main

//https://leetcode-cn.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	temp := l3
	rest := 0
	value := 0
	for l1 != nil && l2 != nil {
		value = (l1.Val + l2.Val + rest) % 10
		rest = (l1.Val + l2.Val + rest) / 10
		temp.Next = &ListNode{Val: value}
		temp = temp.Next
		l1, l2 = l1.Next, l2.Next
	}
	for l1 != nil {
		value = (l1.Val + rest) % 10
		rest = (l1.Val + rest) / 10
		temp.Next = &ListNode{Val: value}
		temp = temp.Next
		l1 = l1.Next
	}
	for l2 != nil {
		value = (l2.Val + rest) % 10
		rest = (l2.Val + rest) / 10
		temp.Next = &ListNode{Val: value}
		temp = temp.Next
		l2 = l2.Next
	}
	if rest != 0 {
		temp.Next = &ListNode{Val: rest}
	}
	return l3.Next
}
