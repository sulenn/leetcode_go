package main

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(T *testing.T) {
	l1 := &ListNode{9, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{9, &ListNode{6, &ListNode{4, nil}}}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
