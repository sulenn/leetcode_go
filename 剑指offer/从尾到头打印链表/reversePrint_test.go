package main

import (
	"fmt"
	"testing"
)

func Test_reversePrint(T *testing.T)  {
	head := &ListNode{1,nil}
	node := head
	node.Next = &ListNode{3,nil}
	node = node.Next
	node.Next = &ListNode{2,nil}
	node = node.Next
	fmt.Println(reversePrint(head))
}
