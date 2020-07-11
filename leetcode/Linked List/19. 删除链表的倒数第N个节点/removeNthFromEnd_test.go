package main

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(T *testing.T) {
	head := &ListNode{5, nil}
	print(removeNthFromEnd(head, 1))
	head.Next = &ListNode{7, nil}
	print(removeNthFromEnd(head, 2))
	print(removeNthFromEnd(head, 1))
	print(removeNthFromEnd(head, 3))
}

func print(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
