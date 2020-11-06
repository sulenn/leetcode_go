package main

//https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res*2 + head.Val
		head = head.Next
	}
	return res
}
