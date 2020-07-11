package main

import (
	"fmt"
	"testing"
)

func Test_reverseKGroup(T *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//head = reverseKGroup(head, 2)
	head = reverseKGroup(head, 3)
	print(head)
}

func print(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
